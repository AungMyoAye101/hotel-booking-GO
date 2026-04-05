package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);unique;not null;index" json:"email" validate:"required,email,max=255"`
	Password  string    `gorm:"not null"`
	City      string    `gorm:"type:varchar(100)" json:"city"`
	Country   string    `gorm:"type:varchar(100)" json:"country"`
	Phone     string    `gorm:"type:varchar(100)" json:"phone"`
	Token     string    `gorm:"type:text" json:"token"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
