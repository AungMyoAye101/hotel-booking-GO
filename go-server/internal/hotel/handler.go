package hotel

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

func (h *Handler) CreateHotel(c echo.Context) error {
	var dto CreateHotelDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hotel, err := h.service.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, "hotel created", hotel)
}

func (h *Handler) GetAllHotels(c echo.Context) error {
	params, err := pagination.Parse(c)
	if err != nil {
		return err
	}

	offset := (params.Page - 1) * params.Limit
	hotels, total, err := h.service.FindAll(offset, params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponseWithMeta(
		c,
		http.StatusOK,
		"hotels fetched",
		hotels,
		pagination.NewMeta(params, total),
	)
}

func (h *Handler) GetHotelByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	hotel, err := h.service.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "hotel not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "hotel fetched", hotel)
}

func (h *Handler) UpdateHotel(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var dto UpdateHotelDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hotel, err := h.service.Update(id, dto)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "hotel not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "hotel updated", hotel)
}

func (h *Handler) DeleteHotel(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := h.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "hotel not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "hotel deleted", nil)
}
