from fastapi import APIRouter
from elasticsearch import Elasticsearch
from .models import SuggestKeywords, SuggestQuery
from .crud import suggest_keyword, suggest_search_query


router = APIRouter(prefix='/suggest', tags=['suggest'])


client = Elasticsearch(
    "http://localhost:9200"
)


@router.get("/query", response_model=SuggestQuery)
async def search_suggest(query: str):
    return suggest_search_query(client, query)


@router.get("/keywords", response_model=SuggestKeywords)
async def suggest_keyword_api(query: str):
    return suggest_keyword(client, query)