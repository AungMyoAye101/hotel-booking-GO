package receipt

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

	api := e.Group("/api/v1/receipt")
	api.POST("", handler.CreateReceipt)
	api.GET("", handler.GetAllReceipts)
	api.GET("/:id", handler.GetReceiptByID)
	api.GET("/user/:userId", handler.GetReceiptsByUserID)
}
