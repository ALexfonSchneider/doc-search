package documents

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"sync"
)

func (r *Repository) GetMetrics(ctx context.Context, wordCloudLimit int) (*entity.Metrics, error) {
	var wordCloud []entity.WordCloudItem
	var years []entity.DocumentsInYearCount

	var err error

	var ws sync.WaitGroup
	ws.Add(2)
	go func() {
		defer ws.Done()
		wordCloud, err = r.GetWordCloud(ctx, wordCloudLimit)
	}()
	go func() {
		defer ws.Done()
		years, err = r.GetYears(ctx)
	}()
	ws.Wait()

	if err != nil {
		return nil, err
	}

	return &entity.Metrics{WordCloud: wordCloud, Years: years}, nil
}
