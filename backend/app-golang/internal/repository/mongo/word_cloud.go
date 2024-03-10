package mongo

import (
	"context"
	"doc-search-app-backend/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetWordCloud(ctx context.Context, limit int) (*[]entities.WordCloudItem, error) {
	pipeline := bson.A{
		bson.D{{"$unwind", bson.D{{"path", "$metrics.word_cloud"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$metrics.word_cloud.value"},
					{"count", bson.D{{"$sum", "$metrics.word_cloud.count"}}},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"count", -1}}}},
		bson.D{{"$limit", limit}},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 0},
					{"value", "$_id"},
					{"count", 1},
				},
			},
		},
	}

	opts := options.Aggregate()
	cursor, err := r.docs.Aggregate(context.Background(), pipeline, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var wordCloudItems []entities.WordCloudItem
	err = cursor.All(ctx, &wordCloudItems)
	if err != nil {
		return nil, err
	}

	return &wordCloudItems, nil
}

func (r *Repository) GetMetrics(ctx context.Context, wordCloudLimit int) (*entities.Metrics, error) {
	wordCloud, err := r.GetWordCloud(ctx, wordCloudLimit)
	if err != nil {
		return nil, err
	}

	return &entities.Metrics{WordCloud: *wordCloud}, nil
}
