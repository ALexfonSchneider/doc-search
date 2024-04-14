package elastic

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"encoding/json"
	"fmt"
	"strings"
)

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
				Article entity.Article `json:"article"`
				Archive entity.Archive `json:"archive"`
				Metrics entity.Metrics `json:"metrics"`
			} `json:"_source"`
			Highlight entity.Highlight
		} `json:"hits"`
	} `json:"hits"`
}

func (r *Repository) SearchArticle(ctx context.Context, query string, keywords []string, year *int, page int, size int) (*entity.SearchResultsPaginate, error) {
	searchQuery := r.BuildSearchQuery(query, keywords, year, page, size)

	fmt.Println(searchQuery)

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

	var articles []entity.Document

	for _, document := range mapRes.Hits.Content {
		articles = append(articles, entity.Document{
			Article:   document.Source.Article,
			Archive:   document.Source.Archive,
			Metrics:   document.Source.Metrics,
			Highlight: &document.Highlight,
		})
	}

	if articles == nil {
		articles = []entity.Document{}
	}

	return &entity.SearchResultsPaginate{
		Articles:  articles,
		Page:      page,
		Size:      size,
		TotalSize: mapRes.Hits.Total.Value,
	}, nil
}

func (r *Repository) GetDocument(ctx context.Context, id string) (*entity.Document, error) {
	response, err := r.TypedClient.Get(r.searchIndex, id).Do(ctx)
	if err != nil {
		return nil, err
	}

	if !response.Found {
		return nil, &entity.DocumentNotFoundErr{Id: id}
	}

	data, err := response.Source_.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var document entity.Document
	if err := json.Unmarshal(data, &document); err != nil {
		return nil, err
	}

	return &document, nil
}
