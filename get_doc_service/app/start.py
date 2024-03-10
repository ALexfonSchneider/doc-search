import json
from src.documents.models import Document
from src.documents.providers.providers import DocumentProvider, DocumentProviderParse
from src.services import GetDocsService

class DocumentProviderFromJson(DocumentProvider):
    def __init__(self, path: str) -> None:
        super().__init__()
        self._path = path
        
    def get_documents(self) -> list[Document]:
        with open(self._path, "r+") as file:
            text = json.loads(file.read())
            documents = [Document(**data) for data in text]
            return documents


if __name__ == "__main__":
    provider = DocumentProviderFromJson("documents.json") 
    
    # with open("documents.json", "w+") as file:
    #     file.write(json.dumps([document.model_dump() for document in provider.get_documents()]))
    
    thread = GetDocsService(provider=provider)
    thread.start()
    thread.join()