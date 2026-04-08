package auth

import (
	"errors"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type Service struct {
	repo          *Repository
	accessSecret  []byte
	refreshSecret []byte
}

func NewService(r *Repository, accessSecret, refreshSecret string) *Service {
	return &Service{
		repo:          r,
		accessSecret:  []byte(accessSecret),
		refreshSecret: []byte(refreshSecret),
	}
}

func (s *Service) Register(dto Register) (*UserResponse, string, error) {
	hashed, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, "", err
	}
	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashed,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, "", err
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Email, "user", s.accessSecret)
	if err != nil {
		return nil, "", err
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String(), user.Email, "user", s.refreshSecret)
	if err != nil {
		return nil, "", err
	}
	user.Token = utils.HashToken(refreshToken)
	if err := s.repo.SaveUser(user); err != nil {
		return nil, "", err
	}

	return &UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Token: accessToken,
	}, refreshToken, nil

}

func (s *Service) Login(dto Login) (*UserResponse, string, error) {
	user, err := s.repo.FindUserByEmail(dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrInvalidCredentials
		}
		return nil, "", err
	}

	ok, upgradedHash, err := utils.VerifyPassword(user.Password, dto.Password)
	if err != nil {
		return nil, "", err
	}
	if !ok {
		return nil, "", ErrInvalidCredentials
	}
	if upgradedHash != "" {
		user.Password = upgradedHash
		if err := s.repo.SaveUser(user); err != nil {
			return nil, "", err
		}
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Email, "user", s.accessSecret)
	if err != nil {
		return nil, "", err
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String(), user.Email, "user", s.refreshSecret)
	if err != nil {
		return nil, "", err
	}
	user.Token = utils.HashToken(refreshToken)
	if err := s.repo.SaveUser(user); err != nil {
		return nil, "", err
	}

	return &UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Token: accessToken,
	}, refreshToken, nil
}

func (s *Service) AdminLogin(dto Login) (*AdminResponse, string, error) {
	admin, err := s.repo.FindAdminByEmail(dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrInvalidCredentials
		}
		return nil, "", err
	}

	ok, upgradedHash, err := utils.VerifyPassword(admin.Password, dto.Password)
	if err != nil {
		return nil, "", err
	}
	if !ok {
		return nil, "", ErrInvalidCredentials
	}
	if upgradedHash != "" {
		admin.Password = upgradedHash
		if err := s.repo.SaveAdmin(admin); err != nil {
			return nil, "", err
		}
	}

	accessToken, err := utils.GenerateAccessToken(admin.ID.String(), admin.Email, admin.Role, s.accessSecret)
	if err != nil {
		return nil, "", err
	}
	refreshToken, err := utils.GenerateRefreshToken(admin.ID.String(), admin.Email, admin.Role, s.refreshSecret)
	if err != nil {
		return nil, "", err
	}
	admin.Token = utils.HashToken(refreshToken)
	if err := s.repo.SaveAdmin(admin); err != nil {
		return nil, "", err
	}

	return &AdminResponse{
		ID:    admin.ID.String(),
		Name:  admin.Name,
		Email: admin.Email,
		Role:  admin.Role,
		Token: accessToken,
	}, refreshToken, nil
}

func (s *Service) Refresh(refreshToken string) (string, string, error) {
	claims, err := utils.ParseToken(refreshToken, s.refreshSecret)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}
	if claims.Typ != "refresh" {
		return "", "", ErrInvalidCredentials
	}

	id, err := uuid.Parse(claims.Sub)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	switch claims.Kind {
	case "admin":
		admin, err := s.repo.FindAdminByID(id)
		if err != nil {
			return "", "", ErrInvalidCredentials
		}
		if admin.Token == "" || !utils.TokensEqualHash(admin.Token, refreshToken) {
			return "", "", ErrInvalidCredentials
		}

		accessToken, err := utils.GenerateAccessToken(admin.ID.String(), admin.Email, admin.Role, s.accessSecret)
		if err != nil {
			return "", "", err
		}
		newRefresh, err := utils.GenerateRefreshToken(admin.ID.String(), admin.Email, admin.Role, s.refreshSecret)
		if err != nil {
			return "", "", err
		}
		admin.Token = utils.HashToken(newRefresh)
		if err := s.repo.SaveAdmin(admin); err != nil {
			return "", "", err
		}
		return accessToken, newRefresh, nil
	default:
		user, err := s.repo.FindUserByID(id)
		if err != nil {
			return "", "", ErrInvalidCredentials
		}
		if user.Token == "" || !utils.TokensEqualHash(user.Token, refreshToken) {
			return "", "", ErrInvalidCredentials
		}

		accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Email, "user", s.accessSecret)
		if err != nil {
			return "", "", err
		}
		newRefresh, err := utils.GenerateRefreshToken(user.ID.String(), user.Email, "user", s.refreshSecret)
		if err != nil {
			return "", "", err
		}
		user.Token = utils.HashToken(newRefresh)
		if err := s.repo.SaveUser(user); err != nil {
			return "", "", err
		}
		return accessToken, newRefresh, nil
	}
}

func (s *Service) CurrentAdmin(refreshToken string) (*AdminResponse, error) {
	claims, err := utils.ParseToken(refreshToken, s.refreshSecret)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if claims.Typ != "refresh" {
		return nil, ErrInvalidCredentials
	}

	id, err := uuid.Parse(claims.Sub)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	admin, err := s.repo.FindAdminByID(id)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if admin.Token == "" || !utils.TokensEqualHash(admin.Token, refreshToken) {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := utils.GenerateAccessToken(admin.ID.String(), admin.Email, admin.Role, s.accessSecret)
	if err != nil {
		return nil, err
	}

	return &AdminResponse{
		ID:    admin.ID.String(),
		Name:  admin.Name,
		Email: admin.Email,
		Role:  admin.Role,
		Token: accessToken,
	}, nil
}
func (s *Service) CurrentUser(refreshToken string) (*UserResponse, error) {
	claims, err := utils.ParseToken(refreshToken, s.refreshSecret)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if claims.Typ != "refresh" {
		return nil, ErrInvalidCredentials
	}
	id, err := uuid.Parse(claims.Sub)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	user, err := s.repo.FindUserByID(id)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if user.Token == "" || !utils.TokensEqualHash(user.Token, refreshToken) {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Email, "user", s.accessSecret)
	if err != nil {
		return nil, err
	}
	newRefresh, err := utils.GenerateRefreshToken(user.ID.String(), user.Email, "user", s.refreshSecret)
	if err != nil {
		return nil, err
	}
	user.Token = utils.HashToken(newRefresh)
	if err := s.repo.SaveUser(user); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Token: accessToken,
	}, nil
}
