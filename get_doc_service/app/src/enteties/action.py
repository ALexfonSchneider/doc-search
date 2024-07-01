import datetime
from pydantic import BaseModel
from .document import Document

STATUS_NEW = "new"
ACTION_ADD = "add"
ACTION_DELETE = "delete"

class Action(BaseModel):
    article_id: str
    status: str
    action: str
    created_at: datetime.datetime
    updated_at: datetime.datetime | None = None
    document: Document | None = None