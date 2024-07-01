package search

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	search Search
}

func NewHandler(search Search) *Handler {
	return &Handler{search: search}
}

func (h *Handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	params := new(Params)
	if err := c.Bind(params); err != nil {
		return err
	}

	go func() {
		err := h.search.IndexQuery(ctx, params.Query)
		fmt.Println(err)
	}()

	result, err := h.search.SearchArticle(ctx, params.Query, params.Keywords, params.Year,
		params.Udk, params.Page, params.Size)
	if err != nil {
		return err
	}

	return c.JSON(200, result)
}
