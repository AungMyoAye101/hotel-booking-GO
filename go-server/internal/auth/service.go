package auth

import "github.com/AungMyoAye101/hotel-booking-GO/pkg/models"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(dto Register) (*models.User, error) {
	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil

}
