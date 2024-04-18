package search_udk

import "doc-search-app-backend/internal/entity"

type SuggestionsResponse struct {
	Suggestions []entity.Udk `json:"suggestions"`
}
