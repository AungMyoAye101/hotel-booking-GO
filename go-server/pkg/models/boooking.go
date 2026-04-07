package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID         uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID      `gorm:"type:uuid;not null" json:"user_id" validate:"required"`
	RoomID     uuid.UUID      `gorm:"type:uuid;not null;index:idx_room_booking" json:"room_id" validate:"required"`
	HotelID    uuid.UUID      `gorm:"type:uuid;not null" json:"hotel_id" validate:"required"`
	Name       string         `gorm:"type:varchar(255)" json:"name" validate:"max=255"`
	Email      string         `gorm:"type:varchar(255)" json:"email" validate:"omitempty,email"`
	City       string         `gorm:"type:varchar(100)" json:"city"`
	Country    string         `gorm:"type:varchar(100)" json:"country"`
	Phone      string         `gorm:"type:varchar(20)" json:"phone"`
	CheckIn    time.Time      `gorm:"not null;index:idx_room_booking;index:idx_checkin_status" json:"check_in" validate:"required"`
	CheckOut   time.Time      `gorm:"not null;index:idx_room_booking;index:idx_checkout_status" json:"check_out" validate:"required,gtfield=CheckIn"`
	Quantity   int            `gorm:"type:int;not null" json:"quantity" validate:"required,gt=0"`
	Guest      int            `gorm:"type:int;not null" json:"guest" validate:"required,gt=0"`
	Status     string         `gorm:"type:varchar(50);index:idx_room_booking;index:idx_checkin_status;index:idx_checkout_status" json:"status" validate:"omitempty,oneof=DRAFT PENDING CONFIRMED STAYED CANCELLED EXPIRED"`
	TotalPrice float64        `gorm:"type:numeric(12,2);not null" json:"total_price" validate:"required,gte=0"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`

	// Associations
	User  *User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"user,omitempty"`
	Room  *Room  `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"room,omitempty"`
	Hotel *Hotel `gorm:"foreignKey:HotelID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"hotel,omitempty"`
}
