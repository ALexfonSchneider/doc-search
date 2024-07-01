import os
import requests
from src.enteties import Archive
from src.parser.parser import parse_article_page, parse_archives_page, parse_archive_page
from uuid import uuid4


def download_article(content: bytearray, folder):            
    file_name = f"{uuid4().hex}.pdf"
    file_path = os.path.join(folder, file_name)
    
    with open(file_path, "x+b") as file:
        file.write(content)
    
    return file_path


def load_article(article_url: str):
    preview_content = requests.get(article_url).text
    article = parse_article_page(preview_content)
    return article


def load_archive(archive: Archive):
    response = requests.get(archive.url)
    content = response.text
    
    archive.articles = [load_article(url) for url in parse_archive_page(content)]
    # archive.articles = [load_article(list(parse_archive_page(content))[0])]
    return archive
     

def load_archives(archives_page_url = "http://journal.asu.ru/psgmm/issue/archive") -> list[Archive]:
    content = requests.get(archives_page_url).text
    archives = [load_archive(archive) for archive in parse_archives_page(content)]
    return archives