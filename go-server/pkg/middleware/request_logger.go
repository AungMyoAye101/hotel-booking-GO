package middleware

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			status := c.Response().Status
			if status == 0 {
				status = 500
			}

			req := c.Request()
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = req.Header.Get(echo.HeaderXRequestID)
			}

			attrs := []any{
				"request_id", requestID,
				"method", req.Method,
				"path", req.URL.Path,
				"query", req.URL.RawQuery,
				"status", status,
				"latency_ms", time.Since(start).Milliseconds(),
				"remote_ip", c.RealIP(),
				"user_agent", req.UserAgent(),
				"bytes_in", req.ContentLength,
				"bytes_out", c.Response().Size,
			}

			if err != nil {
				logger.Error("request", append(attrs, "error", err.Error())...)
			} else {
				logger.Info("request", attrs...)
			}

			return err
		}
	}
}

