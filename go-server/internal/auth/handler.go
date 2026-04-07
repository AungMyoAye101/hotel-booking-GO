package auth

import (
	"net/http"
	"time"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/response"
	"github.com/labstack/echo/v4"
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
	user, refreshToken, err := h.service.Register(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	setRefreshCookie(c, refreshToken)
	return response.SuccessResponse(c, http.StatusCreated, "user registered", user)
}

func (h *Handler) Login(c echo.Context) error {
	var dto Login
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, refreshToken, err := h.service.Login(dto)
	if err != nil {
		if err == ErrInvalidCredentials {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	setRefreshCookie(c, refreshToken)
	return response.SuccessResponse(c, http.StatusOK, "login success", user)
}

func (h *Handler) AdminLogin(c echo.Context) error {
	var dto Login
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	admin, refreshToken, err := h.service.AdminLogin(dto)
	if err != nil {
		if err == ErrInvalidCredentials {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	setRefreshCookie(c, refreshToken)
	return response.SuccessResponse(c, http.StatusOK, "login success", admin)
}

type tokenResponse struct {
	Token string `json:"token"`
}

func (h *Handler) Refresh(c echo.Context) error {
	rc, err := c.Cookie("refresh_token")
	if err != nil || rc == nil || rc.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing refresh token")
	}

	accessToken, newRefresh, err := h.service.Refresh(rc.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid refresh token")
	}
	setRefreshCookie(c, newRefresh)
	return response.SuccessResponse(c, http.StatusOK, "token refreshed", tokenResponse{Token: accessToken})
}

func setRefreshCookie(c echo.Context, refreshToken string) {
	secure := c.Scheme() == "https"
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/api/v1/auth",
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})
}
