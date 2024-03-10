from fastapi import APIRouter, Query, Request
from elasticsearch import Elasticsearch
from .schemes import SearchResultsPaginate
from .crud import index_search, search_articles


router = APIRouter(prefix='/search', tags=['search'])

client = Elasticsearch(
    "http://localhost:9200"
)

# mongo_client = MongoClient('localhost', 27017, username='root', password='password')

def get_elastic_client():
    return client

@router.get("/", response_model=SearchResultsPaginate)
async def search(request: Request, query: str, keywords: list[str] | None = Query([]), page: int = 1, size: int = 10):
    keywords_query = keywords if keywords else request.query_params.getlist("keywords_query[]")
    search_result = search_articles(client, query, keywords_query, page=page, size=size)
    index_search(client, query)
    return search_result