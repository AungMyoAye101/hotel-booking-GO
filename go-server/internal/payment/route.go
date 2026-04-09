package payment

import (
	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/booking"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/receipt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Run(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	_ = cfg

	paymentRepo := NewRepository(db)
	bookingRepo := booking.NewRepository(db)
	receiptRepo := receipt.NewRepository(db)

	bookingService := booking.NewService(bookingRepo)
	receiptService := receipt.NewService(receiptRepo)
	paymentService := NewService(paymentRepo, bookingService, receiptService)

	handler := NewHandler(paymentService)

	api := e.Group("/api/v1/payments")
	api.POST("", handler.CreatePayment)
	api.GET("", handler.GetAllPayments)
	api.GET("/user/:userId", handler.GetPaymentsByUserID)
	api.GET("/:id", handler.GetPaymentByID)
	api.PUT("/:id", handler.UpdatePayment)
	api.DELETE("/:id", handler.DeletePayment)
}
