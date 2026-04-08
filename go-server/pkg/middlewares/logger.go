package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			start := time.Now()
			err := next(c)

			c.Logger().Infof("%s %s %d %v",
				c.Request().Method,
				c.Request().URL.Path,
				c.Response().Status,
				time.Since(start),
			)

			return err
		}
	}
}
