package search_udk

import (
	"context"
	"doc-search-app-backend/internal/entity"
)

type Repository interface {
	SearchUdk(ctx context.Context, query string) ([]entity.Udk, error)
}
