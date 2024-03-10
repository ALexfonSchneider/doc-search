from typing import AsyncGenerator
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.pool import AsyncAdaptedQueuePool
from sqlalchemy.orm import DeclarativeBase

from config import (
    DB_URL,
)

engine = create_async_engine(url=DB_URL, poolclass=AsyncAdaptedQueuePool, isolation_level="REPEATABLE READ")
async_session_maker = sessionmaker(engine, class_=AsyncSession, expire_on_commit=True)


class Base(DeclarativeBase):
    pass


async def get_async_session() -> AsyncGenerator[AsyncSession, None]:
    async with async_session_maker() as session:
        yield session