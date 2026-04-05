package user

import (
	"time"

	"github.com/google/uuid"
)

type ResponseUser struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email" `
	City      string    `json:"city" `
	Country   string    `json:"country"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
