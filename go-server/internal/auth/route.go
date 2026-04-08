package auth

import (
	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo, cfg.AUTH.ACCESS_SECRET, cfg.AUTH.REFRESH_SECRET)
	handler := NewHandler(service)

	api := e.Group("/api/v1/auth")
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
	api.POST("/refresh", handler.Refresh)
	api.GET("/me", handler.CurrentUser)

	admin := api.Group("/admin")
	admin.POST("/login", handler.AdminLogin)
	admin.POST("/refresh", handler.Refresh)
	admin.GET("/me", handler.CurrentAdmin)

}
