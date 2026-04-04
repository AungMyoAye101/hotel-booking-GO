package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database ", err)
	}
	fmt.Println("database connected.")
	return db
}
