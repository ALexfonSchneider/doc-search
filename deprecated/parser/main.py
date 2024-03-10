import asyncio
from serializers import ArchiveSerializer
from loader import load_archives
import json
from tqdm import tqdm
import time


async def main():
    data = await load_archives()
    
    history = {}
    for archive in tqdm(data):
        history[archive.archive_id] = ArchiveSerializer(archive).data
    
    path = r"C:\Users\Alex\Documents\Study\Лекции, пратика\Дипломная\doc-search\elastic\info.json"
    with open(path, "w+") as file:
        file.write(json.dumps(history))


if __name__ == "__main__":
    start = time.time()
    asyncio.run(main())
    print(time.time() - start)