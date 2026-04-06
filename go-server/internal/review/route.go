package review

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

    api := e.Group("/api/v1/reviews")
    api.POST("", handler.CreateReview)
    api.GET("", handler.GetAllReviews)
    api.GET("/:id", handler.GetReviewByID)
    api.PUT("/:id", handler.UpdateReview)
    api.DELETE("/:id", handler.DeleteReview)
}
