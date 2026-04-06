package auth

import (
	"net/http"

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
	user, err := h.service.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, "user registered", user)
}
