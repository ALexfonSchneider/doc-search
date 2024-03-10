import os
from dotenv import load_dotenv

load_dotenv()

#DB
DB_HOST=os.environ.get("DB_HOST")
DB_PORT=os.environ.get("DB_PORT")
DB_USER=os.environ.get("DB_USER")
DB_PASS=os.environ.get("DB_PASSWORD")
DB_NAME=os.environ.get("DB_NAME")

# postgresql://%(DB_USER)s:%(DB_PASS)s@%(DB_HOST)s:%(DB_PORT)s/%(DB_NAME)s
DB_URL=f"postgresql+asyncpg://{DB_USER}:{DB_PASS}@{DB_HOST}:{DB_PORT}/{DB_NAME}"