package elastic

import "fmt"

func (r *Repository) BuildIndexKeyword(query string) string {
	return fmt.Sprintf(r.indexKeywordTemplate, query)
}

func (r *Repository) BuildUnIndexKeyword(query string) string {
	return fmt.Sprintf(r.unIndexKeywordTemplate, query)
}

func (r *Repository) BuildIndexQuery(query string) string {
	return fmt.Sprintf(r.indexQueryQueryTemplate, query)
}

func (r *Repository) BuildSearchQuery(query string, keywords []string, year *int, page int, size int) string {
	var filters string = ""
	var yearQuery string = ""
	var keywordsQuery string = ""

	if year != nil {
		yearQuery = fmt.Sprintf(`{
			"range": {
				"article.published": {
					"gte": "%d",
					"lte": "%d",
					"format": "yyyy"
				}
			}
		}`, *year, (*year)+1)

		filters += yearQuery
	}

	for _, keyword := range keywords {
		if filters != "" {
			filters += ","
		}
		if keywordsQuery != "" {
			keywordsQuery += ","
		}

		filters += fmt.Sprintf(`{"term": {"article.keywords.keyword": "%s"}}`, keyword)
		keywordsQuery += fmt.Sprintf(`{"match": {"article.keywords": {"query": "%s"}}}`, keyword)
	}

	if yearQuery != "" {
		yearQuery += ","
	}

	if keywordsQuery != "" {
		keywordsQuery += ","
	}

	return fmt.Sprintf(r.searchQueryTemplate, size, (page-1)*size, filters, yearQuery, keywordsQuery,
		query, query, query, query, query, query, query)
}

func (r *Repository) BuildSuggestQueriesQuery(query string) string {
	return fmt.Sprintf(r.suggestQueriesQueryTemplate, query)
}

func (r *Repository) BuildSuggestKeywordQuery(query string) string {
	return fmt.Sprintf(r.suggestKeywordQueryTemplate, query)
}
