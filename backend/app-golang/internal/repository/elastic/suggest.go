package elastic

import (
	"context"
	"doc-search-app-backend/internal/entities"
	"encoding/json"
	"fmt"
	"strings"
)

var SuggestKeywordQueryTemplate string
var SuggestQueriesQueryTemplate string

func BuildSuggestKeywordQuery(query string) string {
	return fmt.Sprintf(SuggestKeywordQueryTemplate, query)
}

type keywordsSuggestResult struct {
	Suggest struct {
		KeywordsSuggest []struct {
			Options []struct {
				Content struct {
					KeywordsSuggest struct {
						Value string `json:"input"`
					} `json:"keywords_suggest"`
				} `json:"_source"`
			} `json:"options"`
		} `json:"keywords_suggest"`
	} `json:"suggest"`
}

func (r *Repository) SuggestKeywords(ctx context.Context, query string) (*entities.Suggestions, error) {
	suggestQuery := BuildSuggestKeywordQuery(query)

	fmt.Println(suggestQuery)

	response, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(r.keywordSuggestIndex),
		r.client.Search.WithTrackTotalHits(true),
		r.client.Search.WithBody(strings.NewReader(suggestQuery)),
	)
	if err != nil {
		return nil, err
	}

	var mapRes keywordsSuggestResult
	if err := json.NewDecoder(response.Body).Decode(&mapRes); err != nil {
		return nil, err
	}

	var suggestions []string = make([]string, 0)
	for _, options := range mapRes.Suggest.KeywordsSuggest[0].Options {
		suggestions = append(suggestions, options.Content.KeywordsSuggest.Value)
	}

	return &entities.Suggestions{Suggestions: suggestions}, nil
}

func BuildSuggestQueriesQuery(query string) string {
	return fmt.Sprintf(SuggestQueriesQueryTemplate, query)
}

type queriesSuggestResult struct {
	Suggest struct {
		QuerySuggest []struct {
			Options []struct {
				Content struct {
					QuerySuggest struct {
						Value string `json:"input"`
					} `json:"query"`
				} `json:"_source"`
			} `json:"options"`
		} `json:"queries-suggest"`
	} `json:"suggest"`
}

func (r *Repository) SuggestQueries(ctx context.Context, query string) (*entities.Suggestions, error) {
	suggestQuery := BuildSuggestQueriesQuery(query)

	response, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(r.querySuggestIndex),
		r.client.Search.WithTrackTotalHits(true),
		r.client.Search.WithBody(strings.NewReader(suggestQuery)),
	)
	if err != nil {
		return nil, err
	}

	var mapRes queriesSuggestResult
	if err := json.NewDecoder(response.Body).Decode(&mapRes); err != nil {
		return nil, err
	}

	var suggestions []string = make([]string, 0)
	for _, options := range mapRes.Suggest.QuerySuggest[0].Options {
		suggestions = append(suggestions, options.Content.QuerySuggest.Value)
	}

	return &entities.Suggestions{Suggestions: suggestions}, nil
}