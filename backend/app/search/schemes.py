from pydantic import BaseModel


class Author(BaseModel):
    affiliation: str
    name: str
    orcid: str


class Article(BaseModel):
    # archive_id: str
    article_id: str
    preview_id: str
    title: str
    anatation: str
    keywords: list[str]
    udk: str | None
    published: str
    authors: list[Author]
    link: str | None


class Archive(BaseModel):
    archive_id: str
    name: str
    series: str
    url: str | None


class SearchResult(BaseModel):
    article: Article
    archive: Archive
    metrics: dict
    highlight: dict
    

class SearchResultsPaginate(BaseModel):
    articles: list[SearchResult]
    page: int = 1
    size: int = 10
    total_size: int