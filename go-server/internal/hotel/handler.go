package hotel

import (
	"net/http"
	"strconv"
	"strings"

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

	filter, err := parseHotelFilter(c)
	if err != nil {
		return err
	}

	offset := (params.Page - 1) * params.Limit
	hotels, total, err := h.service.FindAll(offset, params.Limit, filter)
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

func parseHotelFilter(c echo.Context) (HotelFilter, error) {
	filter := HotelFilter{
		Destination: strings.TrimSpace(c.QueryParam("destination")),
		RatingOrder: strings.ToLower(strings.TrimSpace(c.QueryParam("rating_order"))),
		PriceOrder:  strings.ToLower(strings.TrimSpace(c.QueryParam("price_order"))),
	}

	if filter.RatingOrder != "" && filter.RatingOrder != "asc" && filter.RatingOrder != "desc" {
		return filter, echo.NewHTTPError(http.StatusBadRequest, "invalid rating_order; use asc or desc")
	}
	if filter.PriceOrder != "" && filter.PriceOrder != "asc" && filter.PriceOrder != "desc" {
		return filter, echo.NewHTTPError(http.StatusBadRequest, "invalid price_order; use asc or desc")
	}

	if minPriceRaw := strings.TrimSpace(c.QueryParam("min_price")); minPriceRaw != "" {
		minPrice, err := strconv.ParseFloat(minPriceRaw, 64)
		if err != nil || minPrice < 0 {
			return filter, echo.NewHTTPError(http.StatusBadRequest, "invalid min_price")
		}
		filter.MinPrice = minPrice
	}

	if maxPriceRaw := strings.TrimSpace(c.QueryParam("max_price")); maxPriceRaw != "" {
		maxPrice, err := strconv.ParseFloat(maxPriceRaw, 64)
		if err != nil || maxPrice < 0 {
			return filter, echo.NewHTTPError(http.StatusBadRequest, "invalid max_price")
		}
		filter.MaxPrice = maxPrice
	}

	if filter.MinPrice > 0 && filter.MaxPrice > 0 && filter.MaxPrice < filter.MinPrice {
		return filter, echo.NewHTTPError(http.StatusBadRequest, "max_price must be greater than or equal to min_price")
	}

	if starsRaw := strings.TrimSpace(c.QueryParam("stars")); starsRaw != "" {
		for _, raw := range strings.Split(starsRaw, ",") {
			raw = strings.TrimSpace(raw)
			if raw == "" {
				continue
			}
			star, err := strconv.Atoi(raw)
			if err != nil || star < 1 || star > 5 {
				return filter, echo.NewHTTPError(http.StatusBadRequest, "invalid stars list; values must be 1-5")
			}
			filter.Stars = append(filter.Stars, star)
		}
	}

	return filter, nil
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
