package receipt

import "github.com/google/uuid"

type CreateReceiptDTO struct {
	UserID        uuid.UUID `json:"user_id" validate:"required"`
	BookingID     uuid.UUID `json:"booking_id" validate:"required"`
	PaymentID     uuid.UUID `json:"payment_id" validate:"required"`
	PaymentMethod string    `json:"payment_method" validate:"required,oneof=MOBILE_BANKING CARD BANK"`
	Status        string    `json:"status" validate:"required,oneof=PENDING PAID FAILED"`
	Amount        float64   `json:"amount" validate:"required,gt=0"`
}
