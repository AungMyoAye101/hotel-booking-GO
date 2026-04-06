package payment

import (
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
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

