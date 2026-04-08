package user

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

	api := e.Group("/api/v1/users")
	api.Use(middlewares.BearerAuth(cfg.AUTH.ACCESS_SECRET))

	api.GET("", handler.GetAllUsers, middlewares.RequireAdminRoles("admin", "staff"))
	api.GET("/:id", handler.GetUserByID)
	api.PUT("/:id", handler.UpdateUser)
	api.DELETE("/:id", handler.DeleteUser)

}
