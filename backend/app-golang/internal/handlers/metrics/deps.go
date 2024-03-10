package metrics

import (
	"context"
	"doc-search-app-backend/internal/entities"
)

type Metrics interface {
	GetMetrics(ctx context.Context, wordCloudLimit int) (*entities.Metrics, error)
}
