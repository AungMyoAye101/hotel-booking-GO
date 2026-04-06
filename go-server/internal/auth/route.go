package auth

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

	api := e.Group("/api/v1/auth")
	api.POST("/register", handler.Register)
	// api.GET("/login", handler.login)
	// api.GET("/logout", handler.logout)
	// api.PUT("/refresh", handler.Refresh)

}
