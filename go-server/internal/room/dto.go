package room

import "github.com/google/uuid"

type CreateRoomDTO struct {
	Name       string     `json:"name" validate:"required,max=255"`
	MaxPeople  int        `json:"max_people" validate:"required,gt=0"`
	Price      float64    `json:"price" validate:"required,gt=0"`
	TotalRooms int        `json:"total_rooms" validate:"required,gte=0"`
	HotelID    uuid.UUID  `json:"hotel_id" validate:"required"`
	PhotoID    *uuid.UUID `json:"photo_id"`
	BedTypes   string     `json:"bed_types" validate:"required,oneof=king queen full twin single"`
}

type UpdateRoomDTO struct {
	Name       *string    `json:"name" validate:"omitempty,max=255"`
	MaxPeople  *int       `json:"max_people" validate:"omitempty,gt=0"`
	Price      *float64   `json:"price" validate:"omitempty,gt=0"`
	TotalRooms *int       `json:"total_rooms" validate:"omitempty,gte=0"`
	HotelID    *uuid.UUID `json:"hotel_id"`
	PhotoID    *uuid.UUID `json:"photo_id"`
	BedTypes   *string    `json:"bed_types" validate:"omitempty,oneof=king queen full twin single"`
}

