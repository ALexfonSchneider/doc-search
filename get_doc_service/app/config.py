import os
from dotenv import load_dotenv

load_dotenv()

MONGO_URL=os.environ.get("MONGO_URL")