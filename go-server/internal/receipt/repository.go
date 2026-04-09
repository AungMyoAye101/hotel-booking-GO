package receipt

import (
	"errors"
	"fmt"
	"time"

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

func (r *Repository) Create(receipt *models.Receipt) error {
	if err := r.db.Create(receipt).Error; err != nil {
		return err
	}
	// Generate receipt number after ID is set
	now := time.Now()
	receipt.ReceiptNo = fmt.Sprintf("RCPT-%s-%s", now.Format("20060102"), receipt.ID.String()[:8])
	return r.db.Save(receipt).Error
}

func (r *Repository) FindByID(id uuid.UUID) (*models.Receipt, error) {
	var receipt models.Receipt
	if err := r.db.First(&receipt, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &receipt, nil
}

func (r *Repository) FindByUserID(userID uuid.UUID) ([]models.Receipt, error) {
	var receipts []models.Receipt
	if err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&receipts).Error; err != nil {
		return nil, err
	}
	return receipts, nil
}

func (r *Repository) FindAll(offset, limit int) ([]models.Receipt, int64, error) {
	var total int64
	if err := r.db.Model(&models.Receipt{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var receipts []models.Receipt
	if err := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&receipts).Error; err != nil {
		return nil, 0, err
	}
	return receipts, total, nil
}

func (r *Repository) Save(receipt *models.Receipt) error {
	return r.db.Save(receipt).Error
}
