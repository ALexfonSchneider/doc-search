package suggest_keywords

import (
	"context"
	"doc-search-app-backend/internal/entities"
)

type Suggest interface {
	SuggestKeywords(ctx context.Context, query string) (*entities.Suggestions, error)
}
