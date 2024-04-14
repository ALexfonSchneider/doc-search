package cache

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type yearsCache struct {
	Years []entity.DocumentsInYearCount `json:"years"`
}

func (i yearsCache) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (r *Repository) SetYearsToCache(ctx context.Context, years []entity.DocumentsInYearCount) error {
	key := "documents_years"

	cmd := r.client.Set(ctx, key, yearsCache{Years: years}, r.cacheExpireTime)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetYearsFromCache(ctx context.Context) ([]entity.DocumentsInYearCount, error) {
	key := "documents_years"
	cmd := r.client.Get(ctx, key)

	if err := cmd.Err(); err != nil {
		if redis.HasErrorPrefix(err, "redis: nil") {
			return nil, &entity.NoDocumentInCacheErr{}
		}

		return nil, err
	}

	var years yearsCache
	if err := cmd.Scan(&years); err != nil {
		return nil, &entity.NoDocumentInCacheErr{}
	}

	fmt.Println(key, years)

	return years.Years, nil
}
