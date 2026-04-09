package receipt

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

func (h *Handler) CreateReceipt(c echo.Context) error {
	var dto CreateReceiptDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r, err := h.service.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, "receipt created", r)
}

func (h *Handler) GetReceiptsByUserID(c echo.Context) error {
	userID, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	receipts, err := h.service.FindByUserID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "receipts fetched", receipts)
}

func (h *Handler) GetAllReceipts(c echo.Context) error {
	params, err := pagination.Parse(c)
	if err != nil {
		return err
	}

	offset := (params.Page - 1) * params.Limit
	receipts, total, err := h.service.FindAll(offset, params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponseWithMeta(
		c,
		http.StatusOK,
		"receipts fetched",
		receipts,
		pagination.NewMeta(params, total),
	)
}

func (h *Handler) GetReceiptByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	r, err := h.service.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "receipt not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "receipt fetched", r)
}
