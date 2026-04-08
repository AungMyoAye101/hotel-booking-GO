package room

import (
	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := e.Group("/api/v1/rooms")
	api.GET("", handler.GetAllRooms)
	api.GET("/hotel/:hotelId", handler.GetRoomsByHotelID)
	api.GET("/:id", handler.GetRoomByID)

	protected := api.Group("")
	protected.Use(middlewares.BearerAuth(cfg.AUTH.ACCESS_SECRET))
	protected.Use(middlewares.RequireAdminRoles("admin", "staff"))
	protected.POST("", handler.CreateRoom)
	protected.PUT("/:id", handler.UpdateRoom)
	protected.DELETE("/:id", handler.DeleteRoom)
}
