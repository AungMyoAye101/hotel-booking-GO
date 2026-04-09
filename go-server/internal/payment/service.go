package payment

import (
	"github.com/AungMyoAye101/hotel-booking-GO/internal/booking"
	"github.com/AungMyoAye101/hotel-booking-GO/internal/receipt"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo           *Repository
	bookingService *booking.Service
	receiptService *receipt.Service
}

func NewService(r *Repository, bs *booking.Service, rs *receipt.Service) *Service {
	return &Service{
		repo:           r,
		bookingService: bs,
		receiptService: rs,
	}
}

func (s *Service) Create(dto CreatePaymentDTO) (*models.Payment, error) {
	status := dto.Status
	if status == "" {
		status = "PENDING"
	}

	p := &models.Payment{
		BookingID:     dto.BookingID,
		UserID:        dto.UserID,
		PaymentMethod: dto.PaymentMethod,
		Status:        status,
		Amount:        dto.Amount,
	}
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}

	// Update booking status to confirmed
	confirmedStatus := "CONFIRMED"
	_, err := s.bookingService.Update(dto.BookingID, booking.UpdateBookingDTO{
		Status: &confirmedStatus,
	})
	if err != nil {
		return nil, err
	}

	// Create receipt
	_, err = s.receiptService.Create(receipt.CreateReceiptDTO{
		UserID:        dto.UserID,
		BookingID:     dto.BookingID,
		PaymentID:     p.ID,
		PaymentMethod: dto.PaymentMethod,
		Status:        status,
		Amount:        dto.Amount,
	})
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Service) FindAll(offset, limit int) ([]models.Payment, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *Service) FindByID(id uuid.UUID) (*models.Payment, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, dto UpdatePaymentDTO) (*models.Payment, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.PaymentMethod != nil {
		p.PaymentMethod = *dto.PaymentMethod
	}
	if dto.Status != nil {
		p.Status = *dto.Status
	}
	if dto.Amount != nil {
		p.Amount = *dto.Amount
	}

	if err := s.repo.Save(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) Delete(id uuid.UUID) error {
	deleted, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !deleted {
		return gorm.ErrRecordNotFound
	}
	return nil
}
