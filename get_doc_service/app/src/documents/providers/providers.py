
from abc import ABC, abstractmethod
from src.documents.models import Document
from src.documents.parser.loader import load_archives

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