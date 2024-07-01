import datetime
import re
import tempfile
import threading
from time import sleep
import PyPDF2
import requests
from src.providers import DocumentProvider
from src.parser.loader import download_article
from src.enteties import ACTION_ADD, ACTION_DELETE, STATUS_NEW, Action, Document, Metrics
from langdetect import DetectorFactory
import langdetect
from src.logger import logger
import spacy
from spacy.tokens import Token
from src.storage import Storage


email_regex = r"""(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])"""
email_regex_pattern = re.compile(email_regex)


class DocsService:
    def __init__(self, provider: DocumentProvider, storage: Storage, delay=10) -> None:
        self._provider = provider
        self._storage = storage
        self._delay = delay

        DetectorFactory.seed = 0
        
        self.languages = {
            'en': spacy.load('en_core_web_lg'),
            'ru': spacy.load('ru_core_news_lg')
        }
        
        super().__init__()
    
    
    def __word_clean(self, content: str):
        lang = langdetect.detect(content)
        
        if lang not in self.languages.keys():
            raise ValueError('unexpected language')
                    
        doc = self.languages[lang](content)
        
        ents = [ent.text for ent in doc.ents]
                
        words = list(
            filter(
                lambda word: 
                    not word.is_stop 
                    and word.pos_ not in ['SYM', 'SPACE', 'PUNCT', 'PRON', 'PART', 'NUM', 'DET', 'CCONJ', 'CONJ', 'AUX', 'ADV', 'ADP', 'X',      'PROPN']
                    and word.text not in ents
                    # and all([term not in word for term in ['-']])
                    and len(word.lemma_) > 1
                    and not any(x in ['', '', ''] for x in word.lemma_)
                , doc
            )
        )
        
        return words


    def __full_text_clean(self, text: str):
        text = re.sub(r'\\x\d*', '', text)
        text = re.sub(r'\\x00(\d)*', '', text)
        text = text.replace(u"\u0000", "").replace(u"\001f", "").replace(u"\001e", "")
        
        # remove mails
        text = re.sub(email_regex_pattern, '', text)
        
        return text


    def __tokenize(self, text: str):
        text = self.__full_text_clean(text)
        words = self.__word_clean(text)
        
        return words


    def __get_words_cloud(self, tokens: list[Token]):        
        cloud = {}
        for token in tokens:
            word = token.lemma_.lower()
            if word not in cloud:
                cloud[word] = 1
                continue
            cloud[word] += 1
        return [
            {"value": key, "count": count} for key, count in sorted(cloud.items(), key=lambda item: item[1], reverse=True)
        ]
    
    def _database_handle(self, documents: list[Document]):
        income_ids = set(map(lambda x: x.article.article_id, documents))
        exists_ids = set(self._storage.get_documents_ids())
        deleted = exists_ids - income_ids
        new = income_ids - exists_ids
        
        for id in deleted:
            self._storage.add_action(Action(
                article_id=id,
                status=STATUS_NEW,
                action=ACTION_DELETE,
                created_at=datetime.datetime.now(),
            ))
        
        for document in documents:
            if document.article.article_id in new:
                self._storage.add_action(Action(
                    article_id=document.article.article_id,
                    status=STATUS_NEW,
                    action=ACTION_ADD,
                    created_at=datetime.datetime.now(),
                    document=document
                ))
                
        return list(deleted), list(new)
    
    
    def __load_content(self, document: Document):
        response = requests.get(document.article.link)
        with tempfile.TemporaryDirectory() as dir:
            path = download_article(response.content, dir)
            reader = PyPDF2.PdfReader(path)
            content = " ".join([page.extract_text() for page in reader.pages])
            logger.debug(f'!content: {content}')
            match = re.search(r"УДК (?P<udk>\w*.\w*)", content)
            document.article.udk = match.group("udk") if match else None
            document.article.content = content
            
    
    def __calc_metrics(self, document: Document):
        tokens = self.__tokenize(document.article.content)
        
        return Metrics(
            word_cloud=self.__get_words_cloud(tokens)
        )
        
    def as_thead(self) -> threading.Thread:
        return threading.Thread(target=self.start)
                    
    def start(self) -> None:
        logger.debug('start service\n')
        
        while True:
            logger.debug('start collecting documents\n')
            documents = self._provider.get_documents()
            logger.debug(f'get {len(documents)} documents')
            
            for document in documents:
                logger.debug(f'document: {document}')
                self.__load_content(document)
            
                document.metrics = self.__calc_metrics(document)
            
            deleted, new = self._database_handle(documents)
            
            logger.info(f"new: {new}; deleted: {deleted}")
                    
            sleep(self._delay)
            

    