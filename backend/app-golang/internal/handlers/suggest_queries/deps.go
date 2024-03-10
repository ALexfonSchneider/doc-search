package suggest_queries

import (
	"context"
	"doc-search-app-backend/internal/entities"
)

type Suggest interface {
	SuggestQueries(ctx context.Context, query string) (*entities.Suggestions, error)
}
