from elasticsearch import Elasticsearch as elk
from .models import SuggestQuery, SuggestKeywords

def suggest_search_query(client: elk, query: str) -> SuggestQuery:
    response = client.search(suggest={
        "search-suggests": {
            "prefix": query,
            "completion": {    
                "size": 10,        
                "field": "search_suggest"  
            }
        }
    })
    
    result = []
    for suggest in response["suggest"]["search-suggests"][0]["options"]:
        result.append(suggest["_source"]["search_suggest"]["input"])
    
    return SuggestQuery(suggestions=result)


def suggest_keyword(client: elk, query: str) -> SuggestKeywords:
    response = client.search(suggest={
        "keywords_suggest": {
            "prefix": query,
            "completion": {    
                "size": 10,        
                "field": "keywords_suggest"  
            }
        }
    })
    
    keywords = []
    
    for suggest in response["suggest"]["keywords_suggest"][0]["options"]:
        keywords.append(suggest["_source"]["keywords_suggest"]["input"])
    
    return SuggestKeywords(suggestions=keywords)