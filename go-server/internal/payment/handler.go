package payment

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

func (h *Handler) CreatePayment(c echo.Context) error {
	var dto CreatePaymentDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p, err := h.service.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, "payment created", p)
}

func (h *Handler) GetAllPayments(c echo.Context) error {
	params, err := pagination.Parse(c)
	if err != nil {
		return err
	}

	offset := (params.Page - 1) * params.Limit
	payments, total, err := h.service.FindAll(offset, params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponseWithMeta(
		c,
		http.StatusOK,
		"payments fetched",
		payments,
		pagination.NewMeta(params, total),
	)
}

func (h *Handler) GetPaymentByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	p, err := h.service.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "payment not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "payment fetched", p)
}

func (h *Handler) GetPaymentsByUserID(c echo.Context) error {
	userID, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	payments, err := h.service.FindByUserID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "payments fetched", payments)
}

func (h *Handler) UpdatePayment(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var dto UpdatePaymentDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p, err := h.service.Update(id, dto)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "payment not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, "payment updated", p)
}

func (h *Handler) DeletePayment(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := h.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "payment not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "payment deleted", nil)
}
