package elastic

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"encoding/json"
	"strings"
)

func (r *Repository) SearchUdk(ctx context.Context, query string) ([]entity.Udk, error) {
	udkQuery := r.BuildSearchUdkQuery(query)

	response, err := r.TypedClient.Search().Index(r.udkIndex).Raw(strings.NewReader(udkQuery)).Do(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]entity.Udk, 0, 0)
	for _, hit := range response.Hits.Hits {
		data, err := hit.Source_.MarshalJSON()
		if err != nil {
			return nil, err
		}

		var udk entity.Udk
		if err := json.Unmarshal(data, &udk); err != nil {
			return nil, err
		}

		result = append(result, udk)
	}

	return result, nil
}
