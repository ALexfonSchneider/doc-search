import datetime
import hashlib
import re
import tempfile
from threading import Thread
import PyPDF2
import requests
from src.documents.providers.providers import DocumentProvider
from src.documents.parser.loader import download_article
from src.documents.models import Document, Metrics
from pymongo import MongoClient
from langdetect import DetectorFactory
import langdetect
from src.logger import logger
import spacy
from spacy.tokens import Token


email_regex = r"""(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])"""
email_regex_pattern = re.compile(email_regex)


class GetDocsService(Thread):
    def __init__(self, provider: DocumentProvider, mongo_conn: str, delay = 24 * 60 * 60) -> None:
        self._delay = delay
        self._provider = provider
        self._mongo_client = MongoClient(mongo_conn)

        DetectorFactory.seed = 0
        
        self.languages = {
            'en': spacy.load('en_core_web_lg'),
            'ru': spacy.load('ru_core_news_lg')
        }
        
        super().__init__()
    
    
    def __word_clean(self, content: str):
        lang = langdetect.detect(content)
        
        #TODO: обработать
        if lang not in self.languages.keys():
            return []
                    
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
    
    
    def _index_keywords_suggesting(self, keyword: str):
        _id = hashlib.sha256(keyword.encode(encoding='utf-8')).hexdigest()
        self._elastic_client.update(index="keywords-suggest", id=_id, script={
            "source": "ctx._source.keywords_suggest.weight += params.count",
            "params": {
                "count": 1
            }
        },
        upsert={
            "keywords_suggest": {
                "input": keyword,
                "weight": 1
            },
        }, retry_on_conflict=3)
    
    
    def _database_handle(self, documents: list[Document]):
        income_ids = set(map(lambda x: x.article.article_id, documents))
        
        exists = self._mongo_client["doc-search"]["documents"].find({}, {"article.article_id": 1, "hash": 1})
        exists_ids = set(map(lambda x: x["article"]["article_id"], exists))
        
        deleted = exists_ids - income_ids
        new = income_ids - exists_ids
        
        for id in deleted:
            self._mongo_client["doc-search"]["actions"].insert_one({
                "article_id": id,
                "status": "new",
                "action": "delete",
                "created_at": datetime.datetime.now(),
                "updated_at": None
            })
        
        for document in documents:
            if document.article.article_id in new:
                logger.info(f"adding new document {document.article.article_id}")
                document_dict={
                    "article": document.article.model_dump(),
                    "archive": document.archive.model_dump(exclude=["articles"]),
                    "metrics": document.metrics.model_dump()
                }
                self._mongo_client["doc-search"]["actions"].insert_one({
                    "article_id": document.article.article_id,
                    "status": "new",
                    "action": "add",
                    "created_at": datetime.datetime.now(),
                    "updated_at": None,
                    "document": document_dict
                })
                
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

                    
    def run(self) -> None:
        logger.debug('start service\n')
        
        # while True:
        logger.debug('start collecting documents\n')
        documents = self._provider.get_documents()
        logger.debug(f'get {len(documents)} documents')
        
        for document in documents:
            logger.debug(f'document: {document}')
            self.__load_content(document)
        
            document.metrics = self.__calc_metrics(document)
            logger.debug(f'''{document.article.title}: {sorted(document.metrics.word_cloud, key=lambda item: item.count, reverse=True)}\n\n''')
        
        deleted, new = self._database_handle(documents)
        
        logger.info(f"new: {new}; deleted: {deleted}")
                
            # sleep(secs=self._delay)
    