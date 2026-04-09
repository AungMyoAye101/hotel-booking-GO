package payment

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

func (r *Repository) Create(p *models.Payment) error {
	return r.db.Create(p).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.Payment, error) {
	var p models.Payment
	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *Repository) FindByUserID(userID uuid.UUID) ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.
		Preload("User").
		Preload("Booking").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *Repository) FindAll(offset, limit int) ([]models.Payment, int64, error) {
	var total int64
	if err := r.db.Model(&models.Payment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var payments []models.Payment
	if err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&payments).Error; err != nil {
		return nil, 0, err
	}
	return payments, total, nil
}

func (r *Repository) Save(p *models.Payment) error {
	return r.db.Save(p).Error
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	res := r.db.Delete(&models.Payment{}, "id = ?", id)
	return res.RowsAffected > 0, res.Error
}
