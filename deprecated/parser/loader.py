import asyncio
import os
import PyPDF2
import aiohttp
import tempfile
import re

from models import Archive
from page_parser import parse_article_page, parse_archives_page, parse_archive_page
from uuid import NAMESPACE_URL, uuid4


def download_article(content, folder):            
    file_name = f"{uuid4().hex}.pdf"
    file_path = os.path.join(folder, file_name)
    
    with open(file_path, "x+b") as file:
        file.write(content)
    
    return file_path


async def load_article(article_url: str):
    async with aiohttp.ClientSession() as session:
        async with session.get(article_url) as response:
            preview_content = await response.text()  
    article = parse_article_page(preview_content)
    async with aiohttp.ClientSession() as session:
        async with session.get(article.link) as response:
            with tempfile.TemporaryDirectory() as dir:
                path = download_article(await response.read(), dir)
                reader = PyPDF2.PdfReader(path)
                text = " ".join([page.extract_text() for page in reader.pages])
                match = re.search(r"УДК (?P<udk>\w*.\w*)", text)
                article.udk = match.group("udk") if match else None
                print(article.udk)
    print(article.link)
    article.content = text
    return article.clean_text()


async def load_articles(archive: Archive):
    async with aiohttp.ClientSession() as session:
        async with session.get(archive.url) as response:
            content = await response.text()
            tasks = [asyncio.create_task(load_article(article_url)) for article_url in parse_archive_page(content)]
            archive.articles = await asyncio.gather(*tasks)
            return archive
     

async def load_archives(archives_page_url = "http://journal.asu.ru/psgmm/issue/archive") -> list[Archive]:
    async with aiohttp.ClientSession() as session:
        async with session.get(archives_page_url) as response:
            content = await response.text()
            tasks = [asyncio.create_task(load_articles(archive)) for archive in parse_archives_page(content)]
            return await asyncio.gather(*tasks)