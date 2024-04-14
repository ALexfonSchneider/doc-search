package suggest_keywords

import (
	"context"
	"doc-search-app-backend/internal/entity"
)

type Suggest interface {
	SuggestKeywords(ctx context.Context, query string) (*entity.Suggestions, error)
}
