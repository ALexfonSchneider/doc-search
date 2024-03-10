from bs4 import BeautifulSoup, Tag
from ..models import Archive, Article, Author


def parse_archives_page(content: str) -> list[Archive]:
    soup = BeautifulSoup(content, features="lxml")
    for article_container in soup.find_all(attrs={"class": "obj_issue_summary"}):
        title_container: Tag = article_container.find(attrs={"class": "title"})
        series_container: Tag = article_container.find(attrs={"class": "series"})
        
        yield Archive(
            archive_id=title_container["href"].split("/")[-1:][0],
            name=title_container.text,
            series=series_container.text,
        )


def parse_archive_page(content: str) -> list[str]:
    soup = BeautifulSoup(content, features="lxml")
    for section in soup.find_all(attrs={"class": "obj_article_summary"}):
        article_preview_link = section.find(attrs={"class": "title"}).a["href"]
        yield article_preview_link


def parse_article_page(content: str) -> Article:
    soup = BeautifulSoup(content, features="lxml")
    authors = []
    
    try:
        for li in soup.find(attrs={"class": "authors"}).find_all("li"):
            name = li.find(attrs={"class": "name"}).text
            affiliation = li.find(attrs={"class": "affiliation"}).text
            orcid = li.find(attrs={"class": "orcid"}).text.replace("Email:", "")
            authors.append(Author(name=name, affiliation=affiliation, orcid=orcid))      
    except:
        pass 
    
    keywords_container: Tag = soup.find(attrs={"class": "keywords"})
    keywords = [word.strip() for word in keywords_container.find(attrs={"class": "value"}).text.split(",")] if keywords_container else []
        
    published_container = soup.find(attrs={"class": "item published"})
    published = published_container.find(attrs={"class": "value"}).text if published_container else None
    
    title = soup.find(attrs={"class": "page_title"}).text
    anatation = soup.find(attrs={"class": "item abstract"}).find('p')
    anatation = anatation.text if anatation else None
    preview_id, article_id = soup.find(attrs={"class": "obj_galley_link pdf"})["href"].split("/")[-2:]
    archive_id = soup.find(attrs={"class": "item issue"}).find(attrs={"class": "title"})["href"].split("/")[-1:][0]
    
    return Article(
        preview_id=preview_id,
        article_id=article_id,
        title=title,
        keywords=keywords,
        anatation=anatation,
        published=published,
        authors=authors,
        archive_id=archive_id,
    )
    