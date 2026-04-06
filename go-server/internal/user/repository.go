package user

import (
	"errors"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(u *models.User) error {
	return r.db.Create(u).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.User, error) {
	var u models.User
	if err := r.db.First(&u, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *Repository) FindAll(offset, limit int) ([]models.User, int64, error) {
	var total int64
	if err := r.db.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var users []models.User
	if err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *Repository) Save(u *models.User) error {
	return r.db.Save(u).Error
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	res := r.db.Delete(&models.User{}, "id = ?", id)
	return res.RowsAffected > 0, res.Error
}
