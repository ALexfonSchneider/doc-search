package elastic

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

type Repository struct {
	client              *elasticsearch.Client
	TypedClient         *elasticsearch.TypedClient
	searchIndex         string
	keywordSuggestIndex string
	querySuggestIndex   string
	udkIndex            string

	searchQueryTemplate         string
	suggestKeywordQueryTemplate string
	suggestQueriesQueryTemplate string
	indexQueryQueryTemplate     string
	indexKeywordTemplate        string
	unIndexKeywordTemplate      string
	searchUdkTemplate           string
}

func NewRepository(searchIndex string, keywordSuggestIndex string, querySuggestIndex string, udkIndex string, conf elasticsearch.Config) (*Repository, error) {
	client, err := elasticsearch.NewClient(conf)
	if err != nil {
		return nil, err
	}

	TypedClient, err := elasticsearch.NewTypedClient(conf)
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

	IndexKeywordTemplateBytes, err := os.ReadFile("internal/repository/elastic/queries/index-keyword.json")
	if err != nil {
		return nil, err
	}

	UnIndexKeywordTemplateBytes, err := os.ReadFile("internal/repository/elastic/queries/unindex-keyword.json")
	if err != nil {
		return nil, err
	}

	searchUdkTemplateBytes, err := os.ReadFile("internal/repository/elastic/queries/search-udk.json")
	if err != nil {
		return nil, err
	}

	return &Repository{
		client:              client,
		TypedClient:         TypedClient,
		searchIndex:         searchIndex,
		keywordSuggestIndex: keywordSuggestIndex,
		querySuggestIndex:   querySuggestIndex,
		udkIndex:            udkIndex,

		searchQueryTemplate:         string(SearchQueryTemplateBytes),
		suggestKeywordQueryTemplate: string(SuggestKeywordQueryTemplateBytes),
		suggestQueriesQueryTemplate: string(SuggestQueriesTemplateBytes),
		indexQueryQueryTemplate:     string(IndexQueryTemplateBytes),
		indexKeywordTemplate:        string(IndexKeywordTemplateBytes),
		unIndexKeywordTemplate:      string(UnIndexKeywordTemplateBytes),
		searchUdkTemplate:           string(searchUdkTemplateBytes),
	}, nil
}
