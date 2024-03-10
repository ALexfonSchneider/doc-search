package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		if httpError := new(echo.HTTPError); errors.Is(err, httpError.Internal) {
			return err
		}

		fmt.Println(err)

		return echo.ErrInternalServerError
	}
}
