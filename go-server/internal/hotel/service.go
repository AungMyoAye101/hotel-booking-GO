package hotel

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

func (s *Service) Create(dto CreateHotelDTO) (*models.Hotel, error) {
	h := &models.Hotel{
		Name:        dto.Name,
		Description: dto.Description,
		PhotoID:     dto.PhotoID,
		Rating:      dto.Rating,
		Star:        dto.Star,
		Type:        dto.Type,
		Address:     dto.Address,
		Price:       dto.Price,
		City:        dto.City,
		Country:     dto.Country,
		Amenities:   dto.Amenities,
	}
	if err := s.repo.Create(h); err != nil {
		return nil, err
	}
	return h, nil
}

func (s *Service) FindAll(offset, limit int) ([]models.Hotel, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *Service) FindByID(id uuid.UUID) (*models.Hotel, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, dto UpdateHotelDTO) (*models.Hotel, error) {
	h, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		h.Name = *dto.Name
	}
	if dto.Description != nil {
		h.Description = *dto.Description
	}
	if dto.PhotoID != nil {
		h.PhotoID = dto.PhotoID
	}
	if dto.Rating != nil {
		h.Rating = *dto.Rating
	}
	if dto.Star != nil {
		h.Star = *dto.Star
	}
	if dto.Type != nil {
		h.Type = *dto.Type
	}
	if dto.Address != nil {
		h.Address = *dto.Address
	}
	if dto.Price != nil {
		h.Price = *dto.Price
	}
	if dto.City != nil {
		h.City = *dto.City
	}
	if dto.Country != nil {
		h.Country = *dto.Country
	}
	if dto.Amenities != nil {
		h.Amenities = *dto.Amenities
	}

	if err := s.repo.Save(h); err != nil {
		return nil, err
	}
	return h, nil
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
