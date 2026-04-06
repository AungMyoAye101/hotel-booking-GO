package hotel

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

	api := e.Group("/api/v1/hotels")
	api.POST("", handler.CreateHotel)
	api.GET("", handler.GetAllHotels)
	api.GET("/:id", handler.GetHotelByID)
	api.PUT("/:id", handler.UpdateHotel)
	api.DELETE("/:id", handler.DeleteHotel)
}

