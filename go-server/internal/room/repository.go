package room

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

func (r *Repository) Create(room *models.Room) error {
	return r.db.Create(room).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.Room, error) {
	var room models.Room
	if err := r.db.First(&room, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &room, nil
}

func (r *Repository) FindAll(offset, limit int) ([]models.Room, int64, error) {
	var total int64
	if err := r.db.Model(&models.Room{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var rooms []models.Room
	if err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&rooms).Error; err != nil {
		return nil, 0, err
	}

	return rooms, total, nil
}

func (r *Repository) Save(room *models.Room) error {
	return r.db.Save(room).Error
}

func (r *Repository) Delete(id uuid.UUID) (bool, error) {
	res := r.db.Delete(&models.Room{}, "id = ?", id)
	return res.RowsAffected > 0, res.Error
}
