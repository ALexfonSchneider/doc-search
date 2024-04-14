package metrics

import (
	"context"
	"doc-search-app-backend/internal/entity"
)

type Metrics interface {
	GetWordCloud(ctx context.Context, limit int) ([]entity.WordCloudItem, error)
	GetYears(ctx context.Context) ([]entity.DocumentsInYearCount, error)
}

type Cache interface {
	SetWordCloudToCache(ctx context.Context, wordCloud []entity.WordCloudItem, opts entity.CacheOptions) error
	GetWordCloudFromCache(ctx context.Context, opts entity.CacheOptions) ([]entity.WordCloudItem, error)
	SetYearsToCache(ctx context.Context, years []entity.DocumentsInYearCount) error
	GetYearsFromCache(ctx context.Context) ([]entity.DocumentsInYearCount, error)
}
