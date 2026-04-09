package main

import (
	"log"
	"net/http"

	"github.com/AungMyoAye101/hotel-booking-GO/config"

	"github.com/AungMyoAye101/hotel-booking-GO/internal/auth"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/booking"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/hotel"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/payment"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/receipt"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/review"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/room"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/user"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/db"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/response"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal("Failed to load env")
	}
	db, err := db.Connect(cfg.DATABASE.URL)
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	app := NewApp(cfg)
	app.echo.GET("/", func(c echo.Context) error {
		return response.SuccessResponse(c, http.StatusOK, "welcome", "Welcome to Hotel Booking System APIs")
	})

	auth.Run(app.echo, db, cfg)
	user.Run(app.echo, db, cfg)
	hotel.Run(app.echo, db, cfg)
	room.Run(app.echo, db, cfg)
	booking.Run(app.echo, db, cfg)
	review.Run(app.echo, db, cfg)
	payment.Run(app.echo, db, cfg)
	receipt.Run(app.echo, db, cfg)

	app.Start()
}
