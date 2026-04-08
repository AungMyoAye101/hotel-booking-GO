package auth

import (
	"fmt"
	"net/http"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/response"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/utils"
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
	utils.SetRefreshCookie(c, refreshToken)

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
	utils.SetRefreshCookie(c, refreshToken)
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
	utils.SetRefreshCookie(c, refreshToken)
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
	utils.SetRefreshCookie(c, newRefresh)
	return response.SuccessResponse(c, http.StatusOK, "token refreshed", tokenResponse{Token: accessToken})
}

func (h *Handler) CurrentUser(c echo.Context) error {
	rc, err := c.Cookie("refresh_token")
	fmt.Println(rc)
	if err != nil || rc == nil || rc.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing refresh token")
	}
	user, err := h.service.CurrentUser(rc.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	return response.SuccessResponse(c, http.StatusOK, "Get current user", user)
}

func (h *Handler) CurrentAdmin(c echo.Context) error {
	rc, err := c.Cookie("refresh_token")
	if err != nil || rc == nil || rc.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing refresh token")
	}
	user, err := h.service.CurrentAdmin(rc.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	return response.SuccessResponse(c, http.StatusOK, "Get current admin", user)
}
