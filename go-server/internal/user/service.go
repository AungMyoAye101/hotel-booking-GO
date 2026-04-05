package user

import "os/user"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAllUsers() ([]user.User, error) {
	return s.repo.findAll()
}
