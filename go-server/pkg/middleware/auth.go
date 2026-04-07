package middleware

import (
	"net/http"
	"strings"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/utils"
	"github.com/labstack/echo/v4"
)

type Principal struct {
	ID   string
	Kind string // "user" | "admin"
	Role string // for admin/staff
}

const principalCtxKey = "principal"

func GetPrincipal(c echo.Context) (Principal, bool) {
	v := c.Get(principalCtxKey)
	if v == nil {
		return Principal{}, false
	}
	p, ok := v.(Principal)
	return p, ok
}

func BearerAuth(accessSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			if len(auth) < len("Bearer ")+1 || strings.ToLower(auth[:7]) != "bearer " {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header")
			}
			token := strings.TrimSpace(auth[7:])
			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header")
			}

			claims, err := utils.ParseAndVerifyHS256(token, accessSecret)
			if err != nil || claims.Typ != "access" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid access token")
			}

			c.Set(principalCtxKey, Principal{
				ID:   claims.Sub,
				Kind: claims.Aud,
				Role: claims.Role,
			})
			return next(c)
		}
	}
}

func RequireAdminRoles(roles ...string) echo.MiddlewareFunc {
	allowed := map[string]struct{}{}
	for _, r := range roles {
		allowed[r] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			p, ok := GetPrincipal(c)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			if p.Kind != "admin" {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden")
			}
			if _, ok := allowed[p.Role]; !ok {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden")
			}
			return next(c)
		}
	}
}
