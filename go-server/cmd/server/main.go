package main

import (
	"log"
	"net/http"

	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/db"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	database := db.Connect(cfg.DATABASE_URL)
	_ = database
	database.AutoMigrate()
	app := NewApp()

	app.echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome from hotel booking app.")
	})

}
