package hotel

import (
	"errors"
	"strings"

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

func (r *Repository) FindAll(offset, limit int, filter HotelFilter) ([]models.Hotel, int64, error) {
	query := r.db.Model(&models.Hotel{})

	if filter.Destination != "" {
		query = query.Where("LOWER(city) = LOWER(?)", filter.Destination)
	}
	if filter.MinPrice > 0 {
		query = query.Where("price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		query = query.Where("price <= ?", filter.MaxPrice)
	}
	if len(filter.Stars) > 0 {
		query = query.Where("star IN ?", filter.Stars)
	}

	var orderClauses []string
	if filter.RatingOrder != "" {
		orderClauses = append(orderClauses, "rating "+filter.RatingOrder)
	}
	if filter.PriceOrder != "" {
		orderClauses = append(orderClauses, "price "+filter.PriceOrder)
	}
	if len(orderClauses) == 0 {
		orderClauses = append(orderClauses, "created_at desc")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var hotels []models.Hotel
	if err := query.Preload("Photo").Order(strings.Join(orderClauses, ",")).Offset(offset).Limit(limit).Find(&hotels).Error; err != nil {
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
