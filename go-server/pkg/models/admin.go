package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name" validate:"required,max=255"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" validate:"required,email,max=255"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-" validate:"required,min=8"`
	Role      string         `gorm:"type:varchar(50);not null;default:'staff'" json:"role" validate:"required,oneof=admin staff"`
	Token     string         `gorm:"type:text" json:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
