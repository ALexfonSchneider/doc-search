package suggest_queries

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	suggest Suggest
}

func NewHadler(suggest Suggest) *Handler {
	return &Handler{suggest: suggest}
}

func (h *Handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	query := c.QueryParam("query")

	suggestions, err := h.suggest.SuggestQueries(ctx, query)
	if err != nil {
		return err
	}

	c.Request().Header.Set("Content-Type", "application/json")
	return c.JSON(200, suggestions)
}
