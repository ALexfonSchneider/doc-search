package elastic

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

type Repository struct {
	client              *elasticsearch.Client
	searchIndex         string
	keywordSuggestIndex string
	querySuggestIndex   string
}

func NewRepository(searchIndex string, keywordSuggestIndex string, querySuggestIndex string, conf elasticsearch.Config) (*Repository, error) {
	client, err := elasticsearch.NewClient(conf)
	if err != nil {
		return nil, err
	}

	SearchQueryTemplateBytes, err := os.ReadFile("internal/repository/elastic/queries/search.json")
	if err != nil {
		return nil, err
	}

	SuggestKeywordQueryTemplateBytes, err := os.ReadFile("internal/repository/elastic/queries/keywords-suggest.json")
	if err != nil {
		return nil, err
	}

	SuggestQueriesTemplateBytes, err := os.ReadFile("internal/repository/elastic/queries/queries-suggest.json")
	if err != nil {
		return nil, err
	}

	IndexQueryTemplateBytes, err := os.ReadFile("internal/repository/elastic/queries/index-query.json")
	if err != nil {
		return nil, err
	}

	SearchQueryTemplate = string(SearchQueryTemplateBytes)
	SuggestKeywordQueryTemplate = string(SuggestKeywordQueryTemplateBytes)
	SuggestQueriesQueryTemplate = string(SuggestQueriesTemplateBytes)
	IndexQueryQueryTemplate = string(IndexQueryTemplateBytes)

	return &Repository{
		client:              client,
		searchIndex:         searchIndex,
		keywordSuggestIndex: keywordSuggestIndex,
		querySuggestIndex:   querySuggestIndex,
	}, nil
}
