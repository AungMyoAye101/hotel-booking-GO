package utils

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetRefreshCookie(c echo.Context, value string) {
	secure := c.Scheme() == "https"
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func ClearCookie(c echo.Context, name string) {
	secure := c.Scheme() == "https"
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}
