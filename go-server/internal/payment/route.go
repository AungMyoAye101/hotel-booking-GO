package payment

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

	api := e.Group("/api/v1/payments")
	api.POST("", handler.CreatePayment)
	api.GET("", handler.GetAllPayments)
	api.GET("/:id", handler.GetPaymentByID)
	api.PUT("/:id", handler.UpdatePayment)
	api.DELETE("/:id", handler.DeletePayment)
}

