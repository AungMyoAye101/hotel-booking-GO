package booking

import (
	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	_ = cfg

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := e.Group("/api/v1/bookings")
	api.POST("", handler.CreateBooking)
	api.GET("", handler.GetAllBookings)
	api.GET("/:id", handler.GetBookingByID)
	api.PUT("/:id", handler.UpdateBooking)
	api.DELETE("/:id", handler.DeleteBooking)
}

