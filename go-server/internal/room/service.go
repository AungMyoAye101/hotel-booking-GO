package room

import (
	"time"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

type AvailabilityFilter struct {
	CheckIn   time.Time
	CheckOut  time.Time
	MaxPeople *int
}

type AvailableRoom struct {
	models.Room
	AvailableRooms int `json:"available_rooms" gorm:"column:available_rooms"`
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(dto CreateRoomDTO) (*models.Room, error) {
	room := &models.Room{
		Name:       dto.Name,
		MaxPeople:  dto.MaxPeople,
		Price:      dto.Price,
		TotalRooms: dto.TotalRooms,
		HotelID:    dto.HotelID,
		PhotoID:    dto.PhotoID,
		BedTypes:   dto.BedTypes,
	}
	if err := s.repo.Create(room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *Service) FindAll(offset, limit int) ([]models.Room, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *Service) FindAvailableByHotelID(hotelID uuid.UUID, filter AvailabilityFilter) ([]AvailableRoom, error) {
	return s.repo.FindAvailableByHotelID(hotelID, filter)
}

func (s *Service) FindByID(id uuid.UUID) (*models.Room, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, dto UpdateRoomDTO) (*models.Room, error) {
	room, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		room.Name = *dto.Name
	}
	if dto.MaxPeople != nil {
		room.MaxPeople = *dto.MaxPeople
	}
	if dto.Price != nil {
		room.Price = *dto.Price
	}
	if dto.TotalRooms != nil {
		room.TotalRooms = *dto.TotalRooms
	}
	if dto.HotelID != nil {
		room.HotelID = *dto.HotelID
	}
	if dto.PhotoID != nil {
		room.PhotoID = dto.PhotoID
	}
	if dto.BedTypes != nil {
		room.BedTypes = *dto.BedTypes
	}

	if err := s.repo.Save(room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *Service) Delete(id uuid.UUID) error {
	deleted, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !deleted {
		return gorm.ErrRecordNotFound
	}
	return nil
}
