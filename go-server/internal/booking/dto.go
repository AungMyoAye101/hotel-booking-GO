package booking

import (
	"time"

	"github.com/google/uuid"
)

type CreateBookingDTO struct {
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	RoomID     uuid.UUID `json:"room_id" validate:"required"`
	HotelID    uuid.UUID `json:"hotel_id" validate:"required"`
	Name       string    `json:"name" validate:"omitempty,max=255"`
	Email      string    `json:"email" validate:"omitempty,email,max=255"`
	City       string    `json:"city" validate:"omitempty,max=100"`
	Country    string    `json:"country" validate:"omitempty,max=100"`
	Phone      string    `json:"phone" validate:"omitempty,max=20"`
	CheckIn    time.Time `json:"check_in" validate:"required"`
	CheckOut   time.Time `json:"check_out" validate:"required,gtfield=CheckIn"`
	Quantity   int       `json:"quantity" validate:"required,gt=0"`
	Guest      int       `json:"guest" validate:"required,gt=0"`
	Status     string    `json:"status" validate:"omitempty,oneof=DRAFT PENDING CONFIRMED STAYED CANCELLED EXPIRED"`
	TotalPrice float64   `json:"total_price" validate:"required,gte=0"`
}

type UpdateBookingDTO struct {
	Name       *string    `json:"name" validate:"omitempty,max=255"`
	Email      *string    `json:"email" validate:"omitempty,email,max=255"`
	City       *string    `json:"city" validate:"omitempty,max=100"`
	Country    *string    `json:"country" validate:"omitempty,max=100"`
	Phone      *string    `json:"phone" validate:"omitempty,max=20"`
	CheckIn    *time.Time `json:"check_in"`
	CheckOut   *time.Time `json:"check_out"`
	Quantity   *int       `json:"quantity" validate:"omitempty,gt=0"`
	Guest      *int       `json:"guest" validate:"omitempty,gt=0"`
	Status     *string    `json:"status" validate:"omitempty,oneof=DRAFT PENDING CONFIRMED STAYED CANCELLED EXPIRED"`
	TotalPrice *float64   `json:"total_price" validate:"omitempty,gte=0"`
}

