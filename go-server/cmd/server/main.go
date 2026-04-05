package main

import (
	"log"
	"net/http"

	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/user"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/db"
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
		return c.JSON(http.StatusOK, "Welcome to Hotel Booking System APIs")
	})
	user.Run(app.echo, db, cfg)

	app.Start()
}
