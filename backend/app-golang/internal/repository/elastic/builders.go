package elastic

import (
	"fmt"
	"strings"
)

func (r *Repository) BuildIndexKeyword(query string) string {
	return fmt.Sprintf(r.indexKeywordTemplate, query)
}

func (r *Repository) BuildUnIndexKeyword(query string) string {
	return fmt.Sprintf(r.unIndexKeywordTemplate, query)
}

func (r *Repository) BuildIndexQuery(query string) string {
	return fmt.Sprintf(r.indexQueryQueryTemplate, query)
}

// TODO заменить *year на options
func (r *Repository) BuildSearchQuery(query string, keywords []string, year *int, udk *string, page int, size int) string {
	var filters []string
	var queries []string

	if year != nil {
		yearQuery := fmt.Sprintf(`{
			"range": {
				"article.published": {
					"gte": "%d",
					"lte": "%d",
					"format": "yyyy"
				}
			}
		}`, *year, (*year)+1)

		filters = append(filters, yearQuery)
		queries = append(queries, yearQuery)
	}

	if udk != nil {
		filters = append(filters, fmt.Sprintf(`{"term": {"article.udk": "%s"}}`, *udk))
		queries = append(queries, fmt.Sprintf(`{"match": {"article.udk": {"query": "%s"}}}`, *udk))
	}

	for _, keyword := range keywords {
		filters = append(filters, fmt.Sprintf(`{"term": {"article.keywords.keyword": "%s"}}`, keyword))
		queries = append(queries, fmt.Sprintf(`{"match": {"article.keywords": {"query": "%s"}}}`, keyword))
	}

	queriesStr := strings.Join(queries, ",\n")
	if queriesStr != "" {
		queriesStr += ","
	}

	return fmt.Sprintf(r.searchQueryTemplate, size, (page-1)*size,
		strings.Join(filters, ",\n"), queriesStr,
		query, query, query, query, query, query, query)
}

func (r *Repository) BuildSuggestQueriesQuery(query string) string {
	return fmt.Sprintf(r.suggestQueriesQueryTemplate, query)
}

func (r *Repository) BuildSuggestKeywordQuery(query string) string {
	return fmt.Sprintf(r.suggestKeywordQueryTemplate, query)
}

func (r *Repository) BuildSearchUdkQuery(query string) string {
	return fmt.Sprintf(r.searchUdkTemplate, query, query, query)
}
