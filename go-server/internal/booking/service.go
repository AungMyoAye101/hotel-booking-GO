package booking

import (
	"errors"
	"net/http"

	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(dto CreateBookingDTO) (*models.Booking, error) {
	status := "DRAFT"

	b := &models.Booking{
		UserID:     dto.UserID,
		RoomID:     dto.RoomID,
		HotelID:    dto.HotelID,
		Name:       dto.Name,
		Email:      dto.Email,
		City:       dto.City,
		Country:    dto.Country,
		Phone:      dto.Phone,
		CheckIn:    dto.CheckIn,
		CheckOut:   dto.CheckOut,
		Quantity:   dto.Quantity,
		Guest:      dto.Guest,
		Status:     status,
		TotalPrice: dto.TotalPrice,
	}

	if err := s.repo.CreateWithAvailability(b); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, echo.NewHTTPError(http.StatusNotFound, "room not found")
		case errors.Is(err, ErrRoomHotelMismatch):
			return nil, echo.NewHTTPError(http.StatusBadRequest, "room_id does not belong to hotel_id")
		case errors.Is(err, ErrGuestExceedsCapacity):
			return nil, echo.NewHTTPError(http.StatusBadRequest, "guest exceeds room capacity")
		case errors.Is(err, ErrNotEnoughRooms):
			return nil, echo.NewHTTPError(http.StatusConflict, "not enough rooms available for selected dates")
		default:
			return nil, err
		}
	}
	return b, nil
}

func (s *Service) FindAll(offset, limit int) ([]models.Booking, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *Service) FindByUserID(userID uuid.UUID) ([]models.Booking, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) FindByID(id uuid.UUID) (*BookingDetailDTO, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, dto UpdateBookingDTO) (*models.Booking, error) {
	b, err := s.repo.FindModelByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		b.Name = *dto.Name
	}
	if dto.Email != nil {
		b.Email = *dto.Email
	}
	if dto.City != nil {
		b.City = *dto.City
	}
	if dto.Country != nil {
		b.Country = *dto.Country
	}
	if dto.Phone != nil {
		b.Phone = *dto.Phone
	}
	if dto.CheckIn != nil {
		b.CheckIn = *dto.CheckIn
	}
	if dto.CheckOut != nil {
		b.CheckOut = *dto.CheckOut
	}
	if !b.CheckOut.After(b.CheckIn) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "check_out must be after check_in")
	}
	if dto.Quantity != nil {
		b.Quantity = *dto.Quantity
	}
	if dto.Guest != nil {
		b.Guest = *dto.Guest
	}
	if dto.Status != nil {
		b.Status = *dto.Status
	}
	if dto.TotalPrice != nil {
		b.TotalPrice = *dto.TotalPrice
	}

	if err := s.repo.Save(b); err != nil {
		return nil, err
	}
	return b, nil
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
