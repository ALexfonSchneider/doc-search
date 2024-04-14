package cache

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type wordCloudCache struct {
	WordCloud []entity.WordCloudItem `json:"wordCloud"`
}

func (i wordCloudCache) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *wordCloudCache) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

func (r *Repository) SetWordCloudToCache(ctx context.Context, wordCloud []entity.WordCloudItem, opts entity.CacheOptions) error {
	key := fmt.Sprintf("word_cloud_%d", opts.Limit)

	cmd := r.client.Set(ctx, key, wordCloudCache{WordCloud: wordCloud}, r.cacheExpireTime)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetWordCloudFromCache(ctx context.Context, opts entity.CacheOptions) ([]entity.WordCloudItem, error) {
	key := fmt.Sprintf("word_cloud_%d", opts.Limit)
	cmd := r.client.Get(ctx, key)

	if err := cmd.Err(); err != nil {
		if redis.HasErrorPrefix(err, "redis: nil") {
			return nil, &entity.NoDocumentInCacheErr{}
		}

		return nil, err
	}

	var wordCloud wordCloudCache
	if err := cmd.Scan(&wordCloud); err != nil {
		return nil, err
	}

	return wordCloud.WordCloud, nil
}
