package models

import "github.com/google/uuid"

type Hotel struct {
	ID          uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Description string     `gorm:"not null"`
	PhotoID     *uuid.UUID `gorm:"type:uuid"` // Pointer for nullable foreign key
	Rating      float64    `gorm:"index"`
	Star        int        `gorm:"index"`
	Type        string     `gorm:"not null;default:'hotel'"`
	Address     string     `gorm:"not null"`
	Price       float64    `gorm:"not null"`
	City        string     `gorm:json:"city"`
	Country     string
	Amenities   []string `gorm:"serializer:json"`
}
