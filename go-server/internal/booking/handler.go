package booking

import (
	"net/http"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/pagination"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateBooking(c echo.Context) error {
	var dto CreateBookingDTO

	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	b, err := h.service.Create(dto)
	if err != nil {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, "booking created", b)
}

func (h *Handler) GetAllBookings(c echo.Context) error {
	params, err := pagination.Parse(c)
	if err != nil {
		return err
	}

	offset := (params.Page - 1) * params.Limit
	bookings, total, err := h.service.FindAll(offset, params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponseWithMeta(
		c,
		http.StatusOK,
		"bookings fetched",
		bookings,
		pagination.NewMeta(params, total),
	)
}

func (h *Handler) GetBookingByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	b, err := h.service.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "booking not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "booking fetched", b)
}

func (h *Handler) UpdateBooking(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var dto UpdateBookingDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	b, err := h.service.Update(id, dto)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "booking not found")
		}
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return httpErr
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "booking updated", b)
}

func (h *Handler) DeleteBooking(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := h.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "booking not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "booking deleted", nil)
}
