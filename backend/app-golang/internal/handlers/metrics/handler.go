package metrics

import (
	"context"
	"doc-search-app-backend/internal/entity"
	errors2 "errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"strconv"
	"sync"
)

type Handler struct {
	Metrics
	Cache
}

func NewHandler(metrics Metrics, cache Cache) *Handler {
	return &Handler{
		Metrics: metrics,
		Cache:   cache,
	}
}

func (h *Handler) GetMetrics(ctx context.Context, wordCloudLimit int) (*entity.Metrics, error) {
	var wordCloud []entity.WordCloudItem
	var years []entity.DocumentsInYearCount

	var (
		err1 error
		err2 error
	)

	var ws sync.WaitGroup
	ws.Add(2)
	go func() {
		defer ws.Done()

		wordCloud, err1 = h.GetWordCloudFromCache(ctx, entity.CacheOptions{Limit: wordCloudLimit})
		if err1 != nil {
			var err1Inner error

			wordCloud, err1Inner = h.GetWordCloud(ctx, wordCloudLimit)
			if err1Inner != nil {
				err1 = err1Inner
				return
			}
		}

		if errors.Is(err1, &entity.NoDocumentInCacheErr{}) {
			err1 = h.SetWordCloudToCache(ctx, wordCloud, entity.CacheOptions{Limit: wordCloudLimit})
		}
	}()
	go func() {
		defer ws.Done()

		years, err2 = h.GetYearsFromCache(ctx)
		if err2 != nil {
			var err2Inner error

			years, err2Inner = h.GetYears(ctx)
			if err2Inner != nil {
				err2 = err2Inner
				return
			}
		}

		if errors.Is(err2, &entity.NoDocumentInCacheErr{}) {
			err2 = h.SetYearsToCache(ctx, years)
		}
	}()
	ws.Wait()

	metrics := &entity.Metrics{
		WordCloud: wordCloud,
		Years:     years,
	}

	if err := errors2.Join(err1, err2); err != nil {
		return metrics, err
	}

	return metrics, nil
}

func (h *Handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	limitQ := c.QueryParam("limit")

	if limitQ == "" {
		limitQ = "35"
	}

	limit, err := strconv.Atoi(limitQ)
	if err != nil {
		return err
	}

	metrics, err := h.GetMetrics(ctx, limit)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(200, metrics)
}
