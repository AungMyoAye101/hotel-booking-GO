package hotel

import "github.com/google/uuid"

type CreateHotelDTO struct {
	Name        string     `json:"name" validate:"required,max=255"`
	Description string     `json:"description" validate:"required"`
	PhotoID     *uuid.UUID `json:"photo_id"`
	Rating      float64    `json:"rating" validate:"omitempty,min=1,max=10"`
	Star        int        `json:"star" validate:"omitempty,min=1,max=5"`
	Type        string     `json:"type" validate:"required,oneof=hotel motel guest-house"`
	Address     string     `json:"address" validate:"required"`
	Price       float64    `json:"price" validate:"required,gt=0"`
	City        string     `json:"city" validate:"omitempty,max=100"`
	Country     string     `json:"country" validate:"omitempty,max=100"`
	Amenities   []string   `json:"amenities"`
}

type UpdateHotelDTO struct {
	Name        *string    `json:"name" validate:"omitempty,max=255"`
	Description *string    `json:"description" validate:"omitempty"`
	PhotoID     *uuid.UUID `json:"photo_id"`
	Rating      *float64   `json:"rating" validate:"omitempty,min=1,max=10"`
	Star        *int       `json:"star" validate:"omitempty,min=1,max=5"`
	Type        *string    `json:"type" validate:"omitempty,oneof=hotel motel guest-house"`
	Address     *string    `json:"address" validate:"omitempty"`
	Price       *float64   `json:"price" validate:"omitempty,gt=0"`
	City        *string    `json:"city" validate:"omitempty,max=100"`
	Country     *string    `json:"country" validate:"omitempty,max=100"`
	Amenities   *[]string  `json:"amenities"`
}

type HotelFilter struct {
	Destination string  `json:"destination"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	RatingOrder string  `json:"rating_order"`
	PriceOrder  string  `json:"price_order"`
	Stars       []int   `json:"stars"`
}

