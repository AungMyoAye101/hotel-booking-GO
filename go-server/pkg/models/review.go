package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id" validate:"required"`
	HotelID   uuid.UUID      `gorm:"type:uuid;not null" json:"hotel_id" validate:"required"`
	Review    string         `gorm:"type:text;not null" json:"review" validate:"required"`
	Rating    int            `gorm:"type:smallint" json:"rating" validate:"min=1,max=9"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Associations
	User  *User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Hotel *Hotel `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"hotel,omitempty"`
}
