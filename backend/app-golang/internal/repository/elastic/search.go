package elastic

import (
	"context"
	"crypto"
	"doc-search-app-backend/internal/entities"
	"encoding/json"
	"fmt"
	"strings"
)

var SearchQueryTemplate string

func BuildSearchQuery(query string, keywords []string, page int, size int) string {
	var filters string = ""

	for _, keyword := range keywords {
		filters += fmt.Sprintf(`{"term": {"article.keywords.keyword": "%s"}}`, keyword)
	}

	return fmt.Sprintf(SearchQueryTemplate, size, (page-1)*size, filters,
		query, query, query, query, query, query, query)
}

type searchResult struct {
	Hits struct {
		Total struct {
			Relation string `json:"relation"`
			Value    int    `json:"value"`
		} `json:"total"`
		Content []struct {
			FoodId      string  `json:"_id"`
			SearchScore float64 `json:"_score"`
			Source      struct {
				Article entities.Article `json:"article"`
				Archive entities.Archive `json:"archive"`
				Metrics entities.Metrics `json:"metrics"`
			} `json:"_source"`
			Highlight entities.Highlight
		} `json:"hits"`
	} `json:"hits"`
}

func (r *Repository) SearchArticle(ctx context.Context, query string, keywords []string, page int, size int) (*entities.SearchResultsPaginate, error) {
	searchQuery := BuildSearchQuery(query, keywords, page, size)

	response, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(r.searchIndex),
		r.client.Search.WithTrackTotalHits(true),
		r.client.Search.WithBody(strings.NewReader(searchQuery)),
	)
	if err != nil {
		return nil, err
	}

	var mapRes searchResult

	err = json.NewDecoder(response.Body).Decode(&mapRes)
	if err != nil {
		return nil, err
	}

	var articles []entities.SearchResult

	for _, document := range mapRes.Hits.Content {
		articles = append(articles, entities.SearchResult{
			Article:   document.Source.Article,
			Archive:   document.Source.Archive,
			Metrics:   document.Source.Metrics,
			Highlight: document.Highlight,
		})
	}

	if articles == nil {
		articles = []entities.SearchResult{}
	}

	return &entities.SearchResultsPaginate{
		Articles:  articles,
		Page:      page,
		Size:      size,
		TotalSize: mapRes.Hits.Total.Value,
	}, nil
}

var IndexQueryQueryTemplate string

func BuildIndexQuery(query string) string {
	return fmt.Sprintf(IndexQueryQueryTemplate, query)
}

func (r *Repository) IndexQuery(ctx context.Context, query string) error {
	indexQuery := BuildIndexQuery(query)

	h := crypto.SHA256.New()
	h.Write([]byte(query))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	_, err := r.client.Update(r.querySuggestIndex, string(hash), strings.NewReader(indexQuery))
	if err != nil {
		return err
	}

	return nil
}
