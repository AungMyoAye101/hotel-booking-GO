package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Image struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	SecureURL string         `gorm:"type:text;not null" json:"secure_url" validate:"required,url"`
	PublicID  string         `gorm:"type:varchar(255);not null" json:"public_id" validate:"required"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
