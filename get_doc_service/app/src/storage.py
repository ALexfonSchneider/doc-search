from abc import ABC, abstractmethod
from pymongo import MongoClient
from src.enteties import Action


class Storage(ABC):
    @abstractmethod
    def add_action(self, action: Action):
        pass
    
    @abstractmethod
    def get_documents_ids(self) -> list[str]:
        pass
    

class MongoStorage(Storage):
    def __init__(self, mongo_conn: str) -> None:
        self._mongo_client = MongoClient(mongo_conn)

    
    def add_action(self, action: Action):
        document_dict = None
        if action.document:
            document_dict={
                "article": action.document.article.model_dump(),
                "archive": action.document.archive.model_dump(exclude=["articles"]),
                "metrics": action.document.metrics.model_dump()
            }
        
        self._mongo_client["doc-search"]["actions"].insert_one({
            **action.model_dump(),
            "document": document_dict,
        })
        
    def get_documents_ids(self) -> list[str]:
        exists = self._mongo_client["doc-search"]["documents"].find({}, {"article.article_id": 1, "hash": 1})
        return list(map(lambda x: x["article"]["article_id"], exists))
        
        