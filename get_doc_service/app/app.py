from config import MONGO_URL
from src.providers.provider import DocumentProviderJson
from src.services.doc_service.service import DocsService
from src.storage import MongoStorage

if __name__ == "__main__":
    provider = DocumentProviderJson("./documents.json") 
    
    storage = MongoStorage(mongo_conn=MONGO_URL)
    service = DocsService(provider=provider, storage=storage)
    
    thread = service.as_thead()
    thread.start()
    thread.join()