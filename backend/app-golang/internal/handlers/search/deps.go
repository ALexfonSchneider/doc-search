package search

import (
	"context"
	"doc-search-app-backend/internal/entity"
)

type Search interface {
	IndexQuery(ctx context.Context, query string) error
	SearchArticle(ctx context.Context, query string, keywords []string, year *int, page int, size int) (*entity.SearchResultsPaginate, error)
}
