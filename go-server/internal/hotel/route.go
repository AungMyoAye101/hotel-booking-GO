package hotel

import (
	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := e.Group("/api/v1/hotels")
	api.GET("", handler.GetAllHotels)
	api.GET("/:id", handler.GetHotelByID)

	protected := api.Group("")
	protected.Use(middleware.BearerAuth(cfg.AUTH.ACCESS_SECRET))
	protected.Use(middleware.RequireAdminRoles("admin", "staff"))
	protected.POST("", handler.CreateHotel)
	protected.PUT("/:id", handler.UpdateHotel)
	protected.DELETE("/:id", handler.DeleteHotel)
}
