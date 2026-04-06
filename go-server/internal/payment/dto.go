package payment

import "github.com/google/uuid"

type CreatePaymentDTO struct {
	BookingID     uuid.UUID `json:"booking_id" validate:"required"`
	UserID        uuid.UUID `json:"user_id" validate:"required"`
	PaymentMethod string    `json:"payment_method" validate:"required,oneof=MOBILE_BANKING CARD BANK"`
	Status        string    `json:"status" validate:"omitempty,oneof=PENDING PAID FAILED REFUNDED"`
	Amount        float64   `json:"amount" validate:"required,gt=0"`
}

type UpdatePaymentDTO struct {
	PaymentMethod *string  `json:"payment_method" validate:"omitempty,oneof=MOBILE_BANKING CARD BANK"`
	Status        *string  `json:"status" validate:"omitempty,oneof=PENDING PAID FAILED REFUNDED"`
	Amount        *float64 `json:"amount" validate:"omitempty,gt=0"`
}

