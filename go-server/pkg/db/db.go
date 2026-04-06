package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(url string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// if err := db.AutoMigrate(
	// 	&models.Admin{},
	// 	&models.User{},
	// 	&models.Image{},
	// 	&models.Hotel{},
	// 	&models.Room{},
	// 	&models.Booking{},
	// 	&models.Review{},
	// 	&models.Payment{},
	// 	&models.Receipt{},
	// ); err != nil {
	// 	return nil, err
	// }

	return db, nil
}
