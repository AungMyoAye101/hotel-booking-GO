package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name" validate:"required,max=255"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" validate:"required,email,max=255"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-" validate:"required,min=8"`
	City      string         `gorm:"type:varchar(100)" json:"city" validate:"max=100"`
	Country   string         `gorm:"type:varchar(100)" json:"country" validate:"max=100"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone" validate:"max=20"`
	Token     string         `gorm:"type:text" json:"token,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
