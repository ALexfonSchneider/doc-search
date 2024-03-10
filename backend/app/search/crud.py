import hashlib
import json
from elasticsearch import Elasticsearch as elk
from .schemes import SearchResult, SearchResultsPaginate


# {
#     "mappings": {
#       "properties": {
#         "search_suggest": {
#           "type": "completion"
#         }
#       }
#     }
#   }

def index_search(client: elk, query: str):
    _id = hashlib.sha256(query.encode(encoding='utf-8')).hexdigest()
    client.update(index="search-queries", id=_id, script={
        "source": "ctx._source.search_suggest.weight += params.count",
        "params": {
            "count": 1
        },
    }, upsert={
        "search_suggest": {
            "input": query,
            "weight": 1
        },
    }, retry_on_conflict=3)
    

def search_articles(client: elk, query: str, keywords: list[str], page: int, size: int) -> list[SearchResult]:
    query={
            "bool": {
                "filter": [
                    {"term": { "article.keywords.keyword": keyword }} for keyword in keywords
                ],
                "should": [
                    {
                        "match": {
                            "article.authors.name": {
                                "query": query,
                                "boost": 5,
                            }
                        },
                    },
                    {
                        "match": {
                            "article.keywords": {
                                "query": query,
                            }
                        }
                    },
                    {
                        "match": {
                            "article.title": {
                                "query": query,
                                "operator": "and",
                                "boost": 2
                            },
                        }
                    },
                    {
                        "match": {
                            "article.title": {
                                "query": query
                            }
                        }
                    },
                    {
                        "match": {
                            "article.anatation": {
                                "query": query,
                            }
                        }
                    },
                    {
                        "match": {
                            "article.content": {
                                "query": query,
                                "operator": "and",
                                "boost": 2
                            }
                        }
                    },
                    {
                        "match": {
                            "article.content": {
                                "query": query,
                            }
                        }
                    }
                ]
            }
        }
    
    response = client.search(
        index="states",
        from_=(page - 1) * size,
        size=size,
        query=query,
        sort=[
            {"_score": "desc"}   
        ],
        highlight={
            #TODO: расширить подсветку
            "fragment_size":150,
            "fields":{
                "article.content":  { "pre_tags" : ["<b>"], "post_tags" : ["</b>"] },
            }
        }
    )
    articles: list[SearchResult] = []
    
    for item in response.body["hits"]["hits"]:
        articles.append(
            SearchResult(
                **item["_source"],
                highlight=item.get("highlight", {})
            )
        )
        
    total_size = response.body["hits"]["total"]["value"]
    
    return SearchResultsPaginate(
        articles=articles,
        page=page,
        size=size,
        total_size=total_size
    )