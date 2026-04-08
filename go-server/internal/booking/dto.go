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

type BookingUserSummaryDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type BookingHotelSummaryDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Rating   float64   `json:"rating"`
	Star     int       `json:"star"`
	Address  string    `json:"address"`
	PhotoURL string    `json:"photo_url"`
}

type BookingRoomSummaryDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	BedTypes string    `json:"bed_types"`
}

type BookingDetailDTO struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	Phone      string    `json:"phone"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	Quantity   int       `json:"quantity"`
	Guest      int       `json:"guest"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	User  BookingUserSummaryDTO  `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	Hotel BookingHotelSummaryDTO `json:"hotel" gorm:"embedded;embeddedPrefix:hotel_"`
	Room  BookingRoomSummaryDTO  `json:"room" gorm:"embedded;embeddedPrefix:room_"`
}
