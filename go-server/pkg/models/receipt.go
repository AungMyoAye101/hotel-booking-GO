package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Receipt struct {
	ID            uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	ReceiptNo     string         `gorm:"type:varchar(255)" json:"receipt_no"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null;" json:"user_id" validate:"required"`
	BookingID     uuid.UUID      `gorm:"type:uuid;not null" json:"booking_id" validate:"required"`
	PaymentID     uuid.UUID      `gorm:"type:uuid;not null" json:"payment_id" validate:"required"`
	PaymentMethod string         `gorm:"type:varchar(50);not null" json:"payment_method" validate:"required,oneof=MOBILE_BANKING CARD BANK"`
	Status        string         `gorm:"type:varchar(50);not null;default:'PENDING';" json:"status" validate:"required,oneof=PENDING PAID FAILED"`
	Amount        float64        `gorm:"type:numeric(12,2);not null" json:"amount" validate:"required,gt=0"`
	PaidAt        time.Time      `gorm:"autoCreateTime" json:"paid_at"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Associations
	User    *User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"user,omitempty"`
	Booking *Booking `gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"booking,omitempty"`
	Payment *Payment `gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"payment,omitempty"`
}
