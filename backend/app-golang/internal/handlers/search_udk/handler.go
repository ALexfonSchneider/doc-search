package search_udk

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	query := c.QueryParam("query")

	result, err := h.repo.SearchUdk(ctx, query)
	if err != nil {
		return err
	}

	return c.JSON(200, SuggestionsResponse{Suggestions: result})
}
