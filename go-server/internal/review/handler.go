package review

import (
    "net/http"

    "github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
    "github.com/AungMyoAye101/hotel-booking-GO/pkg/pagination"
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

func (h *Handler) CreateReview(c echo.Context) error {
    var dto CreateReviewDTO
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
    return c.JSON(http.StatusCreated, r)
}

func (h *Handler) GetAllReviews(c echo.Context) error {
    params, err := pagination.Parse(c)
    if err != nil {
        return err
    }

    offset := (params.Page - 1) * params.Limit
    reviews, total, err := h.service.FindAll(offset, params.Limit)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, pagination.Response[models.Review]{
        Data: reviews,
        Meta: pagination.NewMeta(params, total),
    })
}

func (h *Handler) GetReviewByID(c echo.Context) error {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
    }

    r, err := h.service.FindByID(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return echo.NewHTTPError(http.StatusNotFound, "review not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, r)
}

func (h *Handler) UpdateReview(c echo.Context) error {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
    }

    var dto UpdateReviewDTO
    if err := c.Bind(&dto); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
    }
    if err := c.Validate(dto); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    r, err := h.service.Update(id, dto)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return echo.NewHTTPError(http.StatusNotFound, "review not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, r)
}

func (h *Handler) DeleteReview(c echo.Context) error {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
    }

    if err := h.service.Delete(id); err != nil {
        if err == gorm.ErrRecordNotFound {
            return echo.NewHTTPError(http.StatusNotFound, "review not found")
        }
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusNoContent)
}
