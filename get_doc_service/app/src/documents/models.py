import re
import string
from typing import Any
from pydantic import BaseModel, computed_field, validator

ARCHIVE_LINK_TEMPLATE = string.Template("http://journal.asu.ru/psgmm/issue/view/$id")
ARTICLE_LINK_TEMPLATE = string.Template("http://journal.asu.ru/psgmm/article/download/$preview_id/$article_id")


def cleanup_text(text):
    if not text:
        return ""
    text = re.sub('(\t+)|(\n+)|(/t*/n+)|(/t+/n*)', ' ', text).strip()
    return text


class Author(BaseModel):
    name: str
    affiliation: str
    orcid: str
    
    def model_post_init(self, __context: Any) -> None:
        self.name = cleanup_text(self.name)
        self.affiliation = cleanup_text(self.affiliation)
        self.orcid = cleanup_text(self.orcid)


class Article(BaseModel):
    preview_id: str
    article_id: str
    title: str = ""
    authors: list[Author] = []
    anatation: str | None = ""
    published: str | None= ""
    keywords: list[str] = [] 
    content: str | None = None
    udk: str | None = None

    @computed_field(return_type=str)
    @property
    def link(self) -> str:
        return ARTICLE_LINK_TEMPLATE.substitute({"preview_id": self.preview_id, "article_id": self.article_id})
    
    def model_post_init(self, __context: Any) -> None:
        self.anatation = cleanup_text(self.anatation)
        self.article_id = cleanup_text(self.article_id)
        self.title = cleanup_text(self.title)
        self.published = cleanup_text(self.published)
        self.preview_id = cleanup_text(self.preview_id)
        self.udk = cleanup_text(self.udk)
        
        [cleanup_text(keyword) for keyword in self.keywords]


class ArchiveInfo(BaseModel):
    archive_id: str
    name: str
    series: str
    
    @computed_field(return_type=str)
    @property
    def url(self):
        return ARCHIVE_LINK_TEMPLATE.substitute({"id": self.archive_id})
    
    def model_post_init(self, __context: Any) -> None:
        self.archive_id = cleanup_text(self.archive_id)
        self.name = cleanup_text(self.name)
        self.series = cleanup_text(self.series)

        
class Archive(ArchiveInfo):
    articles: list[Article] = []


class WordCloudItem(BaseModel):
    value: str
    count: int


class Metrics(BaseModel):
    word_cloud: list[WordCloudItem] = []


class Document(BaseModel):
    archive: ArchiveInfo
    article: Article
    metrics: Metrics
    