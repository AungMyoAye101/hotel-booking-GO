package review

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

func (r *Repository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.Review, error) {
	var review models.Review
	if err := r.db.First(&review, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &review, nil
}

func (r *Repository) FindAll(offset, limit int) ([]models.Review, int64, error) {
	var total int64
	if err := r.db.Model(&models.Review{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var reviews []models.Review
	if err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}

func (r *Repository) FindByHotelID(hotelID uuid.UUID, limit int) ([]models.Review, error) {
	var reviews []models.Review
	if err := r.db.Preload("User").Where("hotel_id = ?", hotelID).Order("created_at desc").Limit(limit).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *Repository) Save(review *models.Review) error {
	return r.db.Save(review).Error
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	res := r.db.Delete(&models.Review{}, "id = ?", id)
	return res.RowsAffected > 0, res.Error
}
