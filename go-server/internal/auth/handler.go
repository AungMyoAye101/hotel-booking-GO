package auth

import (
	"net/http"
	"time"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(c echo.Context) error {
	var dto Register
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid field")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := h.service.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, "user registered", user)
}

type loginResponse struct {
	AccessToken string `json:"accessToken"`
	User        any    `json:"user"`
}

func (h *Handler) Login(c echo.Context) error {
	var dto Login
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, pair, err := h.service.LoginUser(dto)
	if err != nil {
		if err == ErrInvalidCredentials || err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user.Token = ""
	setRefreshCookie(c, userRefreshCookieName, pair.RefreshToken, "/api/v1/auth", 7*24*time.Hour)
	return response.SuccessResponse(c, http.StatusOK, "login success", loginResponse{
		AccessToken: pair.AccessToken,
		User:        user,
	})
}

func (h *Handler) AdminLogin(c echo.Context) error {
	var dto Login
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	admin, pair, err := h.service.LoginAdmin(dto)
	if err != nil {
		if err == ErrInvalidCredentials || err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	admin.Token = ""
	setRefreshCookie(c, adminRefreshCookieName, pair.RefreshToken, "/api/v1/auth/admin", 7*24*time.Hour)
	return response.SuccessResponse(c, http.StatusOK, "login success", loginResponse{
		AccessToken: pair.AccessToken,
		User:        admin,
	})
}

func (h *Handler) Refresh(c echo.Context) error {
	rt, err := c.Cookie(userRefreshCookieName)
	if err != nil || rt.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing refresh token")
	}

	pair, err := h.service.RefreshUser(rt.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid refresh token")
	}

	setRefreshCookie(c, userRefreshCookieName, pair.RefreshToken, "/api/v1/auth", 7*24*time.Hour)
	return response.SuccessResponse(c, http.StatusOK, "token refreshed", map[string]string{
		"accessToken": pair.AccessToken,
	})
}

func (h *Handler) AdminRefresh(c echo.Context) error {
	rt, err := c.Cookie(adminRefreshCookieName)
	if err != nil || rt.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing refresh token")
	}

	pair, err := h.service.RefreshAdmin(rt.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid refresh token")
	}

	setRefreshCookie(c, adminRefreshCookieName, pair.RefreshToken, "/api/v1/auth/admin", 7*24*time.Hour)
	return response.SuccessResponse(c, http.StatusOK, "token refreshed", map[string]string{
		"accessToken": pair.AccessToken,
	})
}

func (h *Handler) Logout(c echo.Context) error {
	var rt string
	if cookie, err := c.Cookie(userRefreshCookieName); err == nil {
		rt = cookie.Value
	}
	_ = h.service.LogoutUser(rt)
	clearCookie(c, userRefreshCookieName, "/api/v1/auth")
	return response.SuccessResponse(c, http.StatusOK, "logged out", nil)
}

func (h *Handler) AdminLogout(c echo.Context) error {
	var rt string
	if cookie, err := c.Cookie(adminRefreshCookieName); err == nil {
		rt = cookie.Value
	}
	_ = h.service.LogoutAdmin(rt)
	clearCookie(c, adminRefreshCookieName, "/api/v1/auth/admin")
	return response.SuccessResponse(c, http.StatusOK, "logged out", nil)
}

const (
	userRefreshCookieName  = "refresh_token"
	adminRefreshCookieName = "refresh_token_admin"
)

func setRefreshCookie(c echo.Context, name, value, path string, ttl time.Duration) {
	secure := c.Scheme() == "https"
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Expires:  time.Now().Add(ttl),
		MaxAge:   int(ttl.Seconds()),
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearCookie(c echo.Context, name, path string) {
	secure := c.Scheme() == "https"
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    "",
		Path:     path,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}
