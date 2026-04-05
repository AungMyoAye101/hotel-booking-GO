package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null;index" json:"name" validate:"required,max=255"`
	Description string         `gorm:"type:text;not null" json:"description" validate:"required"`
	PhotoID     *uuid.UUID     `gorm:"type:uuid" json:"photo_id"`
	Rating      float64        `gorm:"type:numeric(3,1);index" json:"rating" validate:"min=1,max=10"`
	Star        int            `gorm:"type:smallint;index" json:"star" validate:"min=1,max=5"`
	Type        string         `gorm:"type:varchar(50);not null;default:'hotel'" json:"type" validate:"required,oneof=hotel motel guest-house"`
	Address     string         `gorm:"type:text;not null" json:"address" validate:"required"`
	Price       float64        `gorm:"type:numeric(12,2);not null" json:"price" validate:"required,gt=0"`
	City        string         `gorm:"type:varchar(100);index" json:"city"`
	Country     string         `gorm:"type:varchar(100)" json:"country"`
	Amenities   []string       `gorm:"type:jsonb;serializer:json" json:"amenities"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Associations
	Photo *Image `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo,omitempty"`
}
