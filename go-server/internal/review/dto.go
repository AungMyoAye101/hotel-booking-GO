package review

import "github.com/google/uuid"

type CreateReviewDTO struct {
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	HotelID uuid.UUID `json:"hotel_id" validate:"required"`
	Review  string    `json:"review" validate:"required"`
	Rating  int       `json:"rating" validate:"omitempty,min=1,max=9"`
}

type UpdateReviewDTO struct {
	Review *string `json:"review" validate:"omitempty"`
	Rating *int    `json:"rating" validate:"omitempty,min=1,max=9"`
}

