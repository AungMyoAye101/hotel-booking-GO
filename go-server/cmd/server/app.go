package main

import (
	"fmt"
	"log"

	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/middleware"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/response"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/validation"
	"github.com/labstack/echo/v4"
)

type APP struct {
	echo *echo.Echo
	cfg  *config.Config
}

func NewApp(cfg *config.Config) *APP {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Validator = validation.New()
	e.HTTPErrorHandler = response.HTTPErrorHandler
	app := &APP{
		echo: e,
		cfg:  cfg,
	}

	return app
}

func (app *APP) Start() {
	address := fmt.Sprintf(":%s", app.cfg.SERVER.PORT)
	fmt.Println("Server running on port", address)
	if err := app.echo.Start(address); err != nil {
		log.Fatal("Failed to start server")
	}

}
