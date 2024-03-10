import re
import string
from nltk import word_tokenize, pos_tag


STR_OR_NONE = str | None

def cleanup_text(text):
    if not text:
        return ""
    text = re.sub('(\t+)|(\n+)|(/t*/n+)|(/t+/n*)', ' ', text).strip()
    return text

class Author:
    def __init__(self, name, affiliation, orcid) -> None:
        self.name = name
        self.affiliation = affiliation
        self.orcid = orcid
    
    def clean_text(self):
        self.name = cleanup_text(self.name)
        self.affiliation = cleanup_text(self.affiliation)
        self.orcid = cleanup_text(self.orcid)
        
    def __str__(self) -> str:
        return str(self.__dict__)
    
        
class Archive:
    __LINK_TEMPLATE = string.Template("http://journal.asu.ru/psgmm/issue/view/$id")
    
    def __init__(self, archive_id: str, name: str = None, series: str = None) -> None:
        self.archive_id = archive_id
        self.name = name
        self.series = series
        self.articles: list[Article] = []
        
    def clean_text(self):
        self.archive_id = cleanup_text(self.archive_id)
        self.name = cleanup_text(self.name)
        self.series = cleanup_text(self.series)
        
        return self
        
    @property
    def url(self):
        return self.__LINK_TEMPLATE.substitute({"id": self.archive_id})
    
    def __str__(self) -> str:
        return str(self.__dict__)
    
    
class Article:
    __LINK_TEMPLATE = string.Template("http://journal.asu.ru/psgmm/article/download/$preview_id/$article_id")

    
    def __init__(self, preview_id: str, article_id: str, title: str = "", authors: list[Author] = [],
                anatation: str = "", published: str = "",
                archive_id: str = "", keywords: str = "") -> None:
        self.article_id = article_id
        self.authors = authors
        self.title = title
        self.anatation = anatation
        self.published = published
        self.archive_id = archive_id
        self.preview_id = preview_id
        self.keywords = keywords
        self.filename = None
        self.content = None
        self.udk = None
        
    def clean_text(self):
        self.anatation = cleanup_text(self.anatation)
        self.article_id = cleanup_text(self.article_id)
        self.title = cleanup_text(self.title)
        self.published = cleanup_text(self.published)
        self.archive_id = cleanup_text(self.archive_id)
        self.preview_id = cleanup_text(self.preview_id)
        self.udk = cleanup_text(self.udk)
        
        [cleanup_text(keyword) for keyword in self.keywords]
        [author.clean_text() for author in self.authors]
        
        return self
            
        
    @property
    def link(self):
        return self.__LINK_TEMPLATE.substitute({"preview_id": self.preview_id, "article_id": self.article_id})
    
    def __str__(self) -> str:
        return str(self.__dict__)