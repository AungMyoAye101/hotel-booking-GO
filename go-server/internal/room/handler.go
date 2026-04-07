package room

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (h *Handler) CreateRoom(c echo.Context) error {
	var dto CreateRoomDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	room, err := h.service.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, "room created", room)
}

func (h *Handler) GetAllRooms(c echo.Context) error {
	params, err := pagination.Parse(c)
	if err != nil {
		return err
	}

	offset := (params.Page - 1) * params.Limit
	rooms, total, err := h.service.FindAll(offset, params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponseWithMeta(
		c,
		http.StatusOK,
		"rooms fetched",
		rooms,
		pagination.NewMeta(params, total),
	)
}

func (h *Handler) GetRoomsByHotelID(c echo.Context) error {
	hotelID, err := uuid.Parse(c.Param("hotelId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid hotel id")
	}

	checkInRaw := c.QueryParam("check_in")
	if checkInRaw == "" {
		checkInRaw = c.QueryParam("checkin")
	}
	checkOutRaw := c.QueryParam("check_out")
	if checkOutRaw == "" {
		checkOutRaw = c.QueryParam("checkout")
	}
	if checkInRaw == "" || checkOutRaw == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "check_in and check_out are required")
	}

	checkIn, err := parseDateTime(checkInRaw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	checkOut, err := parseDateTime(checkOutRaw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !checkOut.After(checkIn) {
		return echo.NewHTTPError(http.StatusBadRequest, "check_out must be after check_in")
	}

	var maxPeople *int
	maxPeopleRaw := c.QueryParam("max_people")
	if maxPeopleRaw == "" {
		maxPeopleRaw = c.QueryParam("maxPeople")
	}
	if maxPeopleRaw != "" {
		v, err := strconv.Atoi(maxPeopleRaw)
		if err != nil || v <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "max_people must be a positive integer")
		}
		maxPeople = &v
	}

	rooms, err := h.service.FindAvailableByHotelID(hotelID, AvailabilityFilter{
		CheckIn:   checkIn,
		CheckOut:  checkOut,
		MaxPeople: maxPeople,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "available rooms fetched", rooms)
}

func (h *Handler) GetRoomByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	room, err := h.service.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "room not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, "room fetched", room)
}

func (h *Handler) UpdateRoom(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var dto UpdateRoomDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	room, err := h.service.Update(id, dto)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "room not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, "room updated", room)
}

func (h *Handler) DeleteRoom(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := h.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "room not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, "room deleted", nil)
}

func parseDateTime(raw string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339, raw); err == nil {
		return t, nil
	}
	if t, err := time.Parse("2006-01-02", raw); err == nil {
		return t, nil
	}
	return time.Time{}, fmt.Errorf("invalid datetime: %q (use RFC3339 or YYYY-MM-DD)", raw)
}
