package review

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

func (s *Service) Create(dto CreateReviewDTO) (*models.Review, error) {
	r := &models.Review{
		UserID:  dto.UserID,
		HotelID: dto.HotelID,
		Review:  dto.Review,
		Rating:  dto.Rating,
	}
	if err := s.repo.Create(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (s *Service) FindAll(offset, limit int) ([]models.Review, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *Service) FindByID(id uuid.UUID) (*models.Review, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, dto UpdateReviewDTO) (*models.Review, error) {
	r, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Review != nil {
		r.Review = *dto.Review
	}
	if dto.Rating != nil {
		r.Rating = *dto.Rating
	}

	if err := s.repo.Save(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (s *Service) FindByHotelID(hotelID uuid.UUID, limit int) ([]models.Review, error) {
	return s.repo.FindByHotelID(hotelID, limit)
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
