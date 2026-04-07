package user

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

func (s *Service) FindAll(offset, limit int) ([]models.User, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *Service) FindByID(id uuid.UUID) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, dto UpdateUserDTO) (*models.User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		u.Name = *dto.Name
	}
	if dto.Email != nil {
		u.Email = *dto.Email
	}
	if dto.Password != nil {
		u.Password = *dto.Password
	}
	if dto.City != nil {
		u.City = *dto.City
	}
	if dto.Country != nil {
		u.Country = *dto.Country
	}
	if dto.Phone != nil {
		u.Phone = *dto.Phone
	}

	if err := s.repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
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
