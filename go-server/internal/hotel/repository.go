package hotel

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

func (r *Repository) Create(h *models.Hotel) error {
	return r.db.Create(h).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.Hotel, error) {
	var h models.Hotel
	if err := r.db.Preload("Photo").First(&h, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &h, nil
}

func (r *Repository) FindAll(offset, limit int) ([]models.Hotel, int64, error) {
	var total int64
	if err := r.db.Model(&models.Hotel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var hotels []models.Hotel
	if err := r.db.Preload("Photo").Order("created_at desc").Offset(offset).Limit(limit).Find(&hotels).Error; err != nil {
		return nil, 0, err
	}

	return hotels, total, nil
}

func (r *Repository) Save(h *models.Hotel) error {
	return r.db.Save(h).Error
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	res := r.db.Delete(&models.Hotel{}, "id = ?", id)
	return res.RowsAffected > 0, res.Error
}
