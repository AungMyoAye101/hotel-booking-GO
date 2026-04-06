package booking

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

func (r *Repository) Create(b *models.Booking) error {
	return r.db.Create(b).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.Booking, error) {
	var b models.Booking
	if err := r.db.First(&b, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &b, nil
}

func (r *Repository) FindAll(offset, limit int) ([]models.Booking, int64, error) {
	var total int64
	if err := r.db.Model(&models.Booking{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var bookings []models.Booking
	if err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&bookings).Error; err != nil {
		return nil, 0, err
	}

	return bookings, total, nil
}

func (r *Repository) Save(b *models.Booking) error {
	return r.db.Save(b).Error
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	res := r.db.Delete(&models.Booking{}, "id = ?", id)
	return res.RowsAffected > 0, res.Error
}
