package metrics

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type Handler struct {
	Metrics
}

func NewHandler(metrics Metrics) *Handler {
	return &Handler{
		Metrics: metrics,
	}
}

// Handle TODO
func (h *Handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	limit_q := c.QueryParam("limit")

	if limit_q == "" {
		limit_q = "30"
	}

	limit, err := strconv.Atoi(limit_q)
	if err != nil {
		return err
	}

	metrics, err := h.GetMetrics(ctx, limit)
	if err != nil {
		return err
	}

	return c.JSON(200, metrics)
}
