package main

import (
	"fmt"
	"log"

	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APP struct {
	echo *echo.Echo
	cfg  *config.Config
	db   *gorm.DB
}

func NewApp(cfg *config.Config) *APP {
	e := echo.New()

	var gormDB *gorm.DB
	if cfg.DATABASE.Disabled {
		log.Println("Database connection skipped (SKIP_DB=true)")
	} else {
		connectedDB, err := db.Connect(cfg.DATABASE.URL)
		if err != nil {
			log.Fatalf("Failed to connect database: %v", err)
		}
		gormDB = connectedDB
	}
	app := &APP{
		echo: e,
		cfg:  cfg,
		db:   gormDB,
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
