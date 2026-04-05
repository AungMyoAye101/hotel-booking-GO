package db

import (
	"fmt"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`) //Enable the UUID extension in PostgreSQL
	db.AutoMigrate(&models.Admin{}, &models.User{}, &models.Image{}, &models.Hotel{}, &models.Room{}, &models.Booking{}, &models.Review{}, &models.Payment{}, &models.Receipt{})
	fmt.Println("Database sync")
	return db, nil
}
