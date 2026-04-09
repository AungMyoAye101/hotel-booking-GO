package receipt

import (
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(dto CreateReceiptDTO) (*models.Receipt, error) {
	receipt := &models.Receipt{
		UserID:        dto.UserID,
		BookingID:     dto.BookingID,
		PaymentID:     dto.PaymentID,
		PaymentMethod: dto.PaymentMethod,
		Status:        dto.Status,
		Amount:        dto.Amount,
	}
	if err := s.repo.Create(receipt); err != nil {
		return nil, err
	}
	return receipt, nil
}

func (s *Service) FindByID(id uuid.UUID) (*models.Receipt, error) {
	return s.repo.FindByID(id)
}

func (s *Service) FindByUserID(userID uuid.UUID) ([]models.Receipt, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) FindAll(offset, limit int) ([]models.Receipt, int64, error) {
	return s.repo.FindAll(offset, limit)
}
