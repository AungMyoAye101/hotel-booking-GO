package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	ID         uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name       string         `gorm:"type:varchar(255);not null;index" json:"name" validate:"required"`
	MaxPeople  int            `gorm:"type:int;not null" json:"max_people" validate:"required,gt=0"`
	Price      float64        `gorm:"type:numeric(10,2);not null" json:"price" validate:"required,gt=0"`
	TotalRooms int            `gorm:"type:int;not null" json:"total_rooms" validate:"required,gte=0"`
	HotelID    uuid.UUID      `gorm:"type:uuid;not null;" json:"hotel_id" validate:"required"`
	PhotoID    *uuid.UUID     `gorm:"type:uuid" json:"photo_id"`
	BedTypes   string         `gorm:"type:varchar(50);not null" json:"bed_types" validate:"required,oneof=king queen full twin single"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Associations
	Hotel *Hotel `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"hotel,omitempty"`
	Photo *Image `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo,omitempty"`
}
