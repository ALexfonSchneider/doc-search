package suggest_queries

import (
	"context"
	"doc-search-app-backend/internal/entity"
)

type Suggest interface {
	SuggestQueries(ctx context.Context, query string) (*entity.Suggestions, error)
}
