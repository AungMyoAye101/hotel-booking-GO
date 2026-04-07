package booking

import (
	"errors"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

var (
	ErrNotEnoughRooms       = errors.New("not enough rooms available")
	ErrRoomHotelMismatch    = errors.New("room_id does not belong to hotel_id")
	ErrGuestExceedsCapacity = errors.New("guest exceeds room capacity")
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(b *models.Booking) error {
	return r.db.Create(b).Error
}

func (r *Repository) CreateWithAvailability(b *models.Booking) error {
	// Bookings that should not consume inventory.
	excludedStatuses := []string{"CANCELLED", "EXPIRED"}

	return r.db.Transaction(func(tx *gorm.DB) error {
		var room models.Room
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&room, "id = ?", b.RoomID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return gorm.ErrRecordNotFound
			}
			return err
		}

		if room.HotelID != b.HotelID {
			return ErrRoomHotelMismatch
		}
		if b.Guest > (room.MaxPeople * b.Quantity) {
			return ErrGuestExceedsCapacity
		}

		var bookedQty int64
		if err := tx.
			Model(&models.Booking{}).
			Select("COALESCE(SUM(quantity), 0)").
			Where("room_id = ?", b.RoomID).
			Where("status NOT IN ?", excludedStatuses).
			Where("check_in < ? AND check_out > ?", b.CheckOut, b.CheckIn).
			Scan(&bookedQty).Error; err != nil {
			return err
		}

		available := room.TotalRooms - int(bookedQty)
		if available < b.Quantity {
			return ErrNotEnoughRooms
		}

		return tx.Create(b).Error
	})
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
