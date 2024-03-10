from pydantic import BaseModel

class SuggestQuery(BaseModel):
    suggestions: list = []
    

class SuggestKeywords(BaseModel):
    suggestions: list = []