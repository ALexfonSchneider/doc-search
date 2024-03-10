package search

import (
	"context"
	"doc-search-app-backend/internal/entities"
)

type Search interface {
	IndexQuery(ctx context.Context, query string) error
	SearchArticle(ctx context.Context, query string, keywords []string, page int, size int) (*entities.SearchResultsPaginate, error)
}
