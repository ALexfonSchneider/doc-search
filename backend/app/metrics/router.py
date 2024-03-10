from fastapi import APIRouter
from pymongo import MongoClient 

from .crud import get_word_cloud

router = APIRouter(prefix='/metrics', tags=['metrics'])

mongo_client = MongoClient('localhost', 27017, username='root', password='password')

@router.get("/")
def index(count: int = 64):
    return {
        'word_cloud': get_word_cloud(mongo_client, count)
    }