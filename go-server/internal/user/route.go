package user

import (
	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, cfg *config.Config) {

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := e.Group("/api/v1/users")

	api.GET("", handler.GetAllUsers)
	api.GET("/hello", handler.GetUser)

}
