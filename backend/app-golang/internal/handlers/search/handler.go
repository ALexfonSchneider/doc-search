package search

import (
	"context"
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

	err = h.search.IndexQuery(context.Background(), query)

	result, err := h.search.SearchArticle(ctx, query, keywords, page, size)
	if err != nil {
		return err
	}

	return c.JSON(200, result)
}
