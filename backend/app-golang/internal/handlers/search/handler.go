package search

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Handler struct {
	search Search
}

func NewHadler(search Search) *Handler {
	return &Handler{search: search}
}

func (h *Handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	query := c.QueryParam("query")
	keywords, ok := c.QueryParams()["keywords_query[]"]
	if !ok {
		keywords = []string{}
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return err
	}
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		return err
	}

	var Year *int = nil
	yearQ := c.QueryParam("year")
	if yearQ != "" {
		year, err := strconv.Atoi(yearQ)
		if err != nil {
			return err
		}
		Year = &year
	}

	go func() {
		err = h.search.IndexQuery(context.Background(), query)
		fmt.Println(err)
	}()

	result, err := h.search.SearchArticle(ctx, query, keywords, Year, page, size)
	if err != nil {
		return err
	}

	return c.JSON(200, result)
}
