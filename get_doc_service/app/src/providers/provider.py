
from abc import ABC, abstractmethod
import json
from src.enteties import Document
from src.parser.loader import load_archives

class DocumentProvider(ABC):
    @abstractmethod
    def get_documents(self) -> list[Document]:
        pass
    

class DocumentProviderParse(DocumentProvider):
    def get_documents(self) -> list[Document]:
        archives = load_archives()
        
        articles: list[Document] = []
        for archive in archives:
            for article in archive.articles:
                articles.append(
                    Document(archive=archive, article=article)
                )
                
        return articles
    

class DocumentProviderJson(DocumentProvider):
    def __init__(self, path: str) -> None:
        super().__init__()
        self._path = path
        
    def get_documents(self) -> list[Document]:
        with open(self._path, "r+") as file:
            text = json.loads(file.read())
            documents = [Document(**data) for data in text]
            return documents