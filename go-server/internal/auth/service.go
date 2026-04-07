package auth

import (
	"errors"
	"time"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type Service struct {
	repo          *Repository
	accessSecret  string
	refreshSecret string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"-"`
}

func NewService(r *Repository, accessSecret, refreshSecret string) *Service {
	return &Service{
		repo:          r,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     15 * time.Minute,
		refreshTTL:    7 * 24 * time.Hour,
	}
}

func (s *Service) Create(dto Register) (*models.User, error) {
	hashed, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashed,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (s *Service) LoginUser(dto Login) (*models.User, TokenPair, error) {
	var pair TokenPair

	user, err := s.repo.FindUserByEmail(dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pair, ErrInvalidCredentials
		}
		return nil, pair, err
	}

	ok, upgradedHash, err := utils.VerifyPassword(user.Password, dto.Password)
	if err != nil {
		return nil, pair, err
	}
	if !ok {
		return nil, pair, ErrInvalidCredentials
	}
	if upgradedHash != "" {
		user.Password = upgradedHash
		if err := s.repo.SaveUser(user); err != nil {
			return nil, pair, err
		}
	}

	accessToken, refreshToken, err := s.issueTokenPair(user.ID.String(), "user", "user")
	if err != nil {
		return nil, pair, err
	}
	user.Token = utils.SHA256Hex(refreshToken)
	if err := s.repo.SaveUser(user); err != nil {
		return nil, pair, err
	}

	pair.AccessToken = accessToken
	pair.RefreshToken = refreshToken
	return user, pair, nil
}

func (s *Service) LoginAdmin(dto Login) (*models.Admin, TokenPair, error) {
	var pair TokenPair

	admin, err := s.repo.FindAdminByEmail(dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pair, ErrInvalidCredentials
		}
		return nil, pair, err
	}

	ok, upgradedHash, err := utils.VerifyPassword(admin.Password, dto.Password)
	if err != nil {
		return nil, pair, err
	}
	if !ok {
		return nil, pair, ErrInvalidCredentials
	}
	if upgradedHash != "" {
		admin.Password = upgradedHash
		if err := s.repo.SaveAdmin(admin); err != nil {
			return nil, pair, err
		}
	}

	accessToken, refreshToken, err := s.issueTokenPair(admin.ID.String(), "admin", admin.Role)
	if err != nil {
		return nil, pair, err
	}
	admin.Token = utils.SHA256Hex(refreshToken)
	if err := s.repo.SaveAdmin(admin); err != nil {
		return nil, pair, err
	}

	pair.AccessToken = accessToken
	pair.RefreshToken = refreshToken
	return admin, pair, nil
}

func (s *Service) RefreshUser(refreshToken string) (TokenPair, error) {
	return s.refresh(refreshToken, "user")
}

func (s *Service) RefreshAdmin(refreshToken string) (TokenPair, error) {
	return s.refresh(refreshToken, "admin")
}

func (s *Service) LogoutUser(refreshToken string) error {
	return s.logout(refreshToken, "user")
}

func (s *Service) LogoutAdmin(refreshToken string) error {
	return s.logout(refreshToken, "admin")
}

func (s *Service) issueTokenPair(sub, aud, role string) (string, string, error) {
	accessClaims := utils.NewAccessClaims(sub, aud, role, s.accessTTL)
	accessToken, err := utils.SignHS256(accessClaims, s.accessSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := utils.NewRefreshClaims(sub, aud, role, uuid.NewString(), s.refreshTTL)
	refreshToken, err := utils.SignHS256(refreshClaims, s.refreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *Service) refresh(refreshToken string, aud string) (TokenPair, error) {
	var pair TokenPair

	claims, err := utils.ParseAndVerifyHS256(refreshToken, s.refreshSecret)
	if err != nil {
		return pair, err
	}
	if claims.Typ != "refresh" || claims.Aud != aud {
		return pair, utils.ErrInvalidToken
	}

	subID, err := uuid.Parse(claims.Sub)
	if err != nil {
		return pair, utils.ErrInvalidToken
	}

	refreshHash := utils.SHA256Hex(refreshToken)

	if aud == "user" {
		user, err := s.repo.FindUserByID(subID)
		if err != nil {
			return pair, err
		}
		if user.Token == "" || user.Token != refreshHash {
			return pair, utils.ErrInvalidToken
		}

		accessToken, newRefreshToken, err := s.issueTokenPair(user.ID.String(), "user", "user")
		if err != nil {
			return pair, err
		}
		user.Token = utils.SHA256Hex(newRefreshToken)
		if err := s.repo.SaveUser(user); err != nil {
			return pair, err
		}

		pair.AccessToken = accessToken
		pair.RefreshToken = newRefreshToken
		return pair, nil
	}

	admin, err := s.repo.FindAdminByID(subID)
	if err != nil {
		return pair, err
	}
	if admin.Token == "" || admin.Token != refreshHash {
		return pair, utils.ErrInvalidToken
	}

	accessToken, newRefreshToken, err := s.issueTokenPair(admin.ID.String(), "admin", admin.Role)
	if err != nil {
		return pair, err
	}
	admin.Token = utils.SHA256Hex(newRefreshToken)
	if err := s.repo.SaveAdmin(admin); err != nil {
		return pair, err
	}

	pair.AccessToken = accessToken
	pair.RefreshToken = newRefreshToken
	return pair, nil
}

func (s *Service) logout(refreshToken string, aud string) error {
	if refreshToken == "" {
		return nil
	}

	claims, err := utils.ParseAndVerifyHS256(refreshToken, s.refreshSecret)
	if err != nil {
		// If it's expired/invalid, we still consider logout successful from the API perspective.
		return nil
	}
	if claims.Typ != "refresh" || claims.Aud != aud {
		return nil
	}
	subID, err := uuid.Parse(claims.Sub)
	if err != nil {
		return nil
	}

	refreshHash := utils.SHA256Hex(refreshToken)

	if aud == "user" {
		user, err := s.repo.FindUserByID(subID)
		if err != nil {
			return nil
		}
		if user.Token == refreshHash {
			user.Token = ""
			_ = s.repo.SaveUser(user)
		}
		return nil
	}

	admin, err := s.repo.FindAdminByID(subID)
	if err != nil {
		return nil
	}
	if admin.Token == refreshHash {
		admin.Token = ""
		_ = s.repo.SaveAdmin(admin)
	}
	return nil
}
