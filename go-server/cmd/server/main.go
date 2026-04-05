package main

import (
	"log"
	"net/http"

	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal("Failed to load env")
	}
	app := NewApp(cfg)
	app.echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to Hotel Booking System APIs")
	})

	app.Start()
}
