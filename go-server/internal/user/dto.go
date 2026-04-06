package user

import (
	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8"`
	City     string `json:"city" validate:"omitempty,max=100"`
	Country  string `json:"country" validate:"omitempty,max=100"`
	Phone    string `json:"phone" validate:"omitempty,max=20"`
}

type UpdateUserDTO struct {
	Name     *string `json:"name" validate:"omitempty,max=255"`
	Email    *string `json:"email" validate:"omitempty,email,max=255"`
	Password *string `json:"password" validate:"omitempty,min=8"`
	City     *string `json:"city" validate:"omitempty,max=100"`
	Country  *string `json:"country" validate:"omitempty,max=100"`
	Phone    *string `json:"phone" validate:"omitempty,max=20"`
}

type UserIDParam struct {
	ID uuid.UUID `param:"id" validate:"required"`
}
