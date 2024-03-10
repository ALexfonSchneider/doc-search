from pymongo import MongoClient

#TODO: add caching
def get_word_cloud(client: MongoClient, count: int) -> dict:
    collection = client['doc-search']['documents']
    with client.start_session() as session:
        word_cloud = list(collection.aggregate(
            [
                {"$unwind": "$metrics.word_cloud" },
                {"$group": {"_id": "$metrics.word_cloud.value", "count": {"$sum": "$metrics.word_cloud.count"}}},
                {"$sort": {"count": -1}}, {"$limit": count},
                { "$project": {"_id": 0,  "value": "$_id", "count": 1 } }
            ], session=session
        ))
    return word_cloud