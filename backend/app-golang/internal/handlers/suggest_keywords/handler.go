package suggest_keywords

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

	suggestions, err := h.suggest.SuggestKeywords(ctx, query)
	if err != nil {
		return err
	}

	return c.JSON(200, suggestions)
}
