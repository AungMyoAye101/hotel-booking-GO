package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AungMyoAye101/hotel-booking-GO/config"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/db"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/models"
	"github.com/AungMyoAye101/hotel-booking-GO/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	var (
		databaseURL = flag.String("database-url", "", "Postgres connection URL (overrides .env DATABASE_URL when set)")
		migrate     = flag.Bool("migrate", true, "Run GORM AutoMigrate before seeding")
		reset       = flag.Bool("reset", false, "Delete existing rows (hard delete) before seeding")
	)
	flag.Parse()

	url := *databaseURL
	if url == "" {
		cfg, err := config.New()
		if err != nil {
			log.Fatal("Failed to load env (.env): ", err)
		}
		url = cfg.DATABASE.URL
	}

	gormDB, err := db.Connect(url)
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	if err := ensureExtensions(gormDB); err != nil {
		log.Fatal("Failed to ensure DB extensions: ", err)
	}
	if *migrate {
		if err := migrateSchema(gormDB); err != nil {
			log.Fatal("Failed to migrate schema: ", err)
		}
	}

	if *reset {
		if err := resetData(gormDB); err != nil {
			log.Fatal("Failed to reset data: ", err)
		}
	}

	if err := seedAll(gormDB); err != nil {
		log.Fatal("Seeding failed: ", err)
	}

	fmt.Fprintln(os.Stdout, "Seed completed successfully.")
}

func ensureExtensions(db *gorm.DB) error {
	// Needed for uuid_generate_v4() defaults used across models.
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
}

func migrateSchema(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Image{},
		&models.Hotel{},
		&models.Room{},
		&models.Booking{},
		&models.Review{},
		&models.Payment{},
		&models.Receipt{},
	)
}

func resetData(db *gorm.DB) error {
	// Child tables first to satisfy foreign key constraints.
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("1=1").Delete(&models.Receipt{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.Payment{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.Review{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.Booking{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.Room{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.Hotel{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.User{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.Admin{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("1=1").Delete(&models.Image{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func seedAll(db *gorm.DB) error {
	now := time.Now().UTC().Truncate(time.Second)

	//image uuid
	imageHotel1ID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	imageHotel2ID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	imageHotel3ID := uuid.MustParse("33333333-3333-3333-3333-333333333333")
	imageHotel4ID := uuid.MustParse("44444444-4444-4444-4444-444444444444")

	// for room images
	imageRoom1ID := uuid.MustParse("c1c1c1c1-1111-1111-1111-111111111111")
	imageRoom2ID := uuid.MustParse("c2c2c2c2-2222-2222-2222-222222222222")
	imageRoom3ID := uuid.MustParse("c3c3c3c3-3333-3333-3333-333333333333")
	imageRoom4ID := uuid.MustParse("c4c4c4c4-4444-4444-4444-444444444444")

	//hotel id
	hotel1ID := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	hotel2ID := uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")
	hotel3ID := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1")
	hotel4ID := uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb2")
	hotel5ID := uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb3")

	// rooms id
	room1ID := uuid.MustParse("d1d1d1d1-1111-1111-1111-111111111111")
	room2ID := uuid.MustParse("d2d2d2d2-2222-2222-2222-222222222222")
	room3ID := uuid.MustParse("d3d3d3d3-3333-3333-3333-333333333333")
	room4ID := uuid.MustParse("d4d4d4d4-4444-4444-4444-444444444444")
	room5ID := uuid.MustParse("d5d5d5d5-5555-5555-5555-555555555555")
	room6ID := uuid.MustParse("d6d6d6d6-6666-6666-6666-666666666666")
	room7ID := uuid.MustParse("d7d7d7d7-7777-7777-7777-777777777777")
	room8ID := uuid.MustParse("d8d8d8d8-8888-8888-8888-888888888888")

	//users id
	user1ID := uuid.MustParse("eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee")
	user2ID := uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")
	admin1ID := uuid.MustParse("99999999-9999-9999-9999-999999999999")
	//booking id
	booking1ID := uuid.MustParse("12121212-1212-1212-1212-121212121212")
	booking2ID := uuid.MustParse("34343434-3434-3434-3434-343434343434")
	//payment id
	payment1ID := uuid.MustParse("56565656-5656-5656-5656-565656565656")
	payment2ID := uuid.MustParse("78787878-7878-7878-7878-787878787878")
	//receive id
	receipt1ID := uuid.MustParse("90909090-9090-9090-9090-909090909090")
	review1ID := uuid.MustParse("abababab-abab-abab-abab-abababababab")
	review2ID := uuid.MustParse("cdcdcdcd-cdcd-cdcd-cdcd-cdcdcdcdcdcd")

	hashedUserPass, err := utils.HashPassword("password123")
	if err != nil {
		return err
	}
	hashedAdminPass, err := utils.HashPassword("admin123")
	if err != nil {
		return err
	}

	seedRefreshHash := utils.HashToken("seed-refresh-token")

	images := []models.Image{
		{
			ID:        imageHotel1ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310527/Booking/wsw8p37h1civp44an6fi.webp",
			PublicID:  "hotel_1",
			CreatedAt: now.Add(-72 * time.Hour),
			UpdatedAt: now.Add(-72 * time.Hour),
		},
		{
			ID:        imageHotel2ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310527/Booking/wsw8p37h1civp44an6fi.webp",
			PublicID:  "hotel_2",
			CreatedAt: now.Add(-48 * time.Hour),
			UpdatedAt: now.Add(-48 * time.Hour),
		},
		{
			ID:        imageHotel3ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310499/Booking/i5wzesnjyurymzxcgxtr.webp",
			PublicID:  "hotel_3",
			CreatedAt: now.Add(-72 * time.Hour),
			UpdatedAt: now.Add(-72 * time.Hour),
		},
		{
			ID:        imageHotel4ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310519/Booking/moohn7uc2w8vtvib8ooh.webp",
			PublicID:  "hotel_4",
			CreatedAt: now.Add(-48 * time.Hour),
			UpdatedAt: now.Add(-48 * time.Hour),
		},
		{
			ID:        imageRoom1ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310687/Booking/ycgpkyitp8ytqtlnbyax.webp",
			PublicID:  "room_1",
			CreatedAt: now.Add(-48 * time.Hour),
			UpdatedAt: now.Add(-48 * time.Hour),
		},
		{
			ID:        imageRoom2ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310685/Booking/suwgltlf19qrmnbwvvk9.webp",
			PublicID:  "room_2",
			CreatedAt: now.Add(-48 * time.Hour),
			UpdatedAt: now.Add(-48 * time.Hour),
		},
		{
			ID:        imageRoom3ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310684/Booking/ld593btcvg0n9xx0b8ne.webp",
			PublicID:  "room_3",
			CreatedAt: now.Add(-48 * time.Hour),
			UpdatedAt: now.Add(-48 * time.Hour),
		},
		{
			ID:        imageRoom4ID,
			SecureURL: "https://res.cloudinary.com/dnxnmdcjb/image/upload/v1771310683/Booking/iq22krm2sdgdnfyl8l5d.webp",
			PublicID:  "room_4",
			CreatedAt: now.Add(-48 * time.Hour),
			UpdatedAt: now.Add(-48 * time.Hour),
		},
	}

	hotel1PhotoID := imageHotel1ID
	hotel2PhotoID := imageHotel2ID
	hotel3PhotoID := imageHotel3ID
	hotel4PhotoID := imageHotel4ID

	hotels := []models.Hotel{
		{
			ID:          hotel1ID,
			Name:        "Golden Shwedagon Hotel",
			Description: "A luxurious colonial-style hotel located in the heart of downtown Yangon, offering stunning views of the pagoda.",
			PhotoID:     &hotel1PhotoID,
			Rating:      9.5,
			Star:        5,
			Type:        "hotel",
			Address:     "92 Strand Rd, Downtown",
			Price:       150.00,
			City:        "Yangon",
			Country:     "Myanmar",
			Amenities:   []string{"wifi", "parking", "pool", "ac", "gym", "breakfast"},
		},
		{
			ID:          hotel2ID,
			Name:        "Royal Mandalay Haven",
			Description: "Cozy and affordable stay near the Mandalay Palace. Famous for its traditional Mohinga breakfast.",
			PhotoID:     &hotel2PhotoID,
			Rating:      8.2,
			Star:        2,
			Type:        "guest-house",
			Address:     "No 45, 27th Street, Chan Aye Tharzan Township",
			Price:       25.50,
			City:        "Mandalay",
			Country:     "Myanmar",
			Amenities:   []string{"wifi", "parking", "pool", "ac", "gym", "breakfast"},
		},
		{
			ID:          hotel3ID,
			Name:        "Bagan Sunrise Rest",
			Description: "Conveniently located just a short e-bike ride from the ancient temples. Perfect for backpackers and early risers.",
			PhotoID:     &hotel3PhotoID,
			Rating:      7.8,
			Star:        3,
			Type:        "motel",
			Address:     "Anawrahta Road, Nyaung-U",
			Price:       40.00,
			City:        "Bagan",
			Country:     "Myanmar",
			Amenities:   []string{"wifi", "parking", "pool", "ac", "gym", "breakfast"},
		},
		{
			ID:          hotel4ID,
			Name:        "Lotus Inle Resort",
			Description: "Beautiful over-water bungalows with breathtaking views of the lake and the Shan hills.",
			PhotoID:     &hotel4PhotoID,
			Rating:      9.1,
			Star:        4,
			Type:        "hotel",
			Address:     "Khaung Daing Village",
			Price:       110.00,
			City:        "Nyaungshwe",
			Country:     "Myanmar",
			Amenities:   []string{"wifi", "parking", "pool", "ac", "gym", "breakfast"},
		},
		{
			ID:          hotel5ID,
			Name:        "Capital Oasis Business Stay",
			Description: "Modern accommodations catering to business travelers with large conference rooms and high-speed internet.",
			PhotoID:     &hotel1PhotoID,
			Rating:      8.5,
			Star:        4,
			Type:        "hotel",
			Address:     "Hotel Zone, Dekkhina Thiri Township",
			Price:       85.00,
			City:        "Naypyidaw",
			Country:     "Myanmar",
			Amenities:   []string{"wifi", "parking", "pool", "ac", "gym", "breakfast"},
		},
	}

	room1PhotoID := imageRoom1ID
	room2PhotoID := imageRoom2ID
	room3PhotoID := imageRoom3ID
	room4PhotoID := imageRoom4ID

	rooms := []models.Room{
		{
			ID:         room1ID,
			Name:       "Deluxe King Room",
			MaxPeople:  2,
			Price:      129.00,
			TotalRooms: 10,
			HotelID:    hotel1ID,
			PhotoID:    &room1PhotoID,
			BedTypes:   "king",
			CreatedAt:  now.Add(-70 * time.Hour),
			UpdatedAt:  now.Add(-50 * time.Hour),
		},
		{
			ID:         room2ID,
			Name:       "Family Queen Room",
			MaxPeople:  4,
			Price:      159.00,
			TotalRooms: 6,
			HotelID:    hotel1ID,
			PhotoID:    &room2PhotoID,
			BedTypes:   "queen",
			CreatedAt:  now.Add(-65 * time.Hour),
			UpdatedAt:  now.Add(-49 * time.Hour),
		},
		{
			ID:         room3ID, // Fixed Duplicate ID
			Name:       "Deluxe King Room",
			MaxPeople:  2,
			Price:      129.00,
			TotalRooms: 10,
			HotelID:    hotel2ID,
			PhotoID:    &room3PhotoID, // Updated image reference
			BedTypes:   "king",
			CreatedAt:  now.Add(-70 * time.Hour),
			UpdatedAt:  now.Add(-50 * time.Hour),
		},
		{
			ID:         room4ID, // Fixed Duplicate ID
			Name:       "Family Queen Room",
			MaxPeople:  4,
			Price:      159.00,
			TotalRooms: 6,
			HotelID:    hotel2ID,
			PhotoID:    &room4PhotoID, // Updated image reference
			BedTypes:   "queen",
			CreatedAt:  now.Add(-65 * time.Hour),
			UpdatedAt:  now.Add(-49 * time.Hour),
		},
		{
			ID:         room5ID, // Fixed Duplicate ID
			Name:       "Deluxe King Room",
			MaxPeople:  2,
			Price:      129.00,
			TotalRooms: 10,
			HotelID:    hotel3ID,
			PhotoID:    &room1PhotoID,
			BedTypes:   "king",
			CreatedAt:  now.Add(-70 * time.Hour),
			UpdatedAt:  now.Add(-50 * time.Hour),
		},
		{
			ID:         room6ID, // Fixed Duplicate ID
			Name:       "Family Queen Room",
			MaxPeople:  4,
			Price:      159.00,
			TotalRooms: 6,
			HotelID:    hotel4ID,
			PhotoID:    &room2PhotoID,
			BedTypes:   "queen",
			CreatedAt:  now.Add(-65 * time.Hour),
			UpdatedAt:  now.Add(-49 * time.Hour),
		},
		{
			ID:         room7ID, // Fixed Duplicate ID
			Name:       "Deluxe King Room",
			MaxPeople:  2,
			Price:      129.00,
			TotalRooms: 10,
			HotelID:    hotel5ID,
			PhotoID:    &room1PhotoID,
			BedTypes:   "king",
			CreatedAt:  now.Add(-70 * time.Hour),
			UpdatedAt:  now.Add(-50 * time.Hour),
		},
		{
			ID:         room8ID, // Fixed Duplicate ID
			Name:       "Family Queen Room",
			MaxPeople:  4,
			Price:      159.00,
			TotalRooms: 6,
			HotelID:    hotel3ID,
			PhotoID:    &room2PhotoID,
			BedTypes:   "queen",
			CreatedAt:  now.Add(-65 * time.Hour),
			UpdatedAt:  now.Add(-49 * time.Hour),
		},
	}

	users := []models.User{
		{
			ID:        user1ID,
			Name:      "Tester",
			Email:     "tester001@gmail.com",
			Password:  hashedUserPass,
			City:      "Yangon",
			Country:   "Myanmar",
			Phone:     "+95-900000001",
			Token:     seedRefreshHash,
			CreatedAt: now.Add(-72 * time.Hour),
			UpdatedAt: now.Add(-2 * time.Hour),
		},
		{
			ID:        user2ID,
			Name:      "Tester 2",
			Email:     "tester002@gmail.com",
			Password:  hashedUserPass,
			City:      "Mandalay",
			Country:   "Myanmar",
			Phone:     "+95-900000002",
			Token:     seedRefreshHash,
			CreatedAt: now.Add(-60 * time.Hour),
			UpdatedAt: now.Add(-3 * time.Hour),
		},
	}

	admins := []models.Admin{
		{
			ID:        admin1ID,
			Name:      "Admin",
			Email:     "admin@gmail.com",
			Password:  hashedAdminPass,
			Role:      "admin",
			Token:     seedRefreshHash,
			CreatedAt: now.Add(-72 * time.Hour),
			UpdatedAt: now.Add(-2 * time.Hour),
		},
	}

	checkIn1 := now.Add(24 * time.Hour)
	checkOut1 := now.Add(48 * time.Hour)
	checkIn2 := now.Add(72 * time.Hour)
	checkOut2 := now.Add(96 * time.Hour)

	bookings := []models.Booking{
		{
			ID:         booking1ID,
			UserID:     user1ID,
			RoomID:     room1ID,
			HotelID:    hotel1ID,
			Name:       "Tester 2",
			Email:      "tester002@gmail.com",
			City:       "Yangon",
			Country:    "Myanmar",
			Phone:      "+95-900000001",
			CheckIn:    checkIn1,
			CheckOut:   checkOut1,
			Quantity:   1,
			Guest:      2,
			Status:     "CONFIRMED",
			TotalPrice: 129.00,
			CreatedAt:  now.Add(-10 * time.Hour),
			UpdatedAt:  now.Add(-9 * time.Hour),
		},
		{
			ID:         booking2ID,
			UserID:     user2ID,
			RoomID:     room2ID,
			HotelID:    hotel1ID,
			Name:       "Tester 1",
			Email:      "tester001@gmail.com",
			City:       "Mandalay",
			Country:    "Myanmar",
			Phone:      "+95-900000002",
			CheckIn:    checkIn2,
			CheckOut:   checkOut2,
			Quantity:   1,
			Guest:      3,
			Status:     "PENDING",
			TotalPrice: 159.00,
			CreatedAt:  now.Add(-8 * time.Hour),
			UpdatedAt:  now.Add(-7 * time.Hour),
		},
	}

	payments := []models.Payment{
		{
			ID:            payment1ID,
			BookingID:     booking1ID,
			UserID:        user1ID,
			PaymentMethod: "CARD",
			Status:        "PAID",
			Amount:        129.00,
			PaidAt:        now.Add(-6 * time.Hour),
			CreatedAt:     now.Add(-6 * time.Hour),
			UpdatedAt:     now.Add(-5 * time.Hour),
		},
		{
			ID:            payment2ID,
			BookingID:     booking2ID,
			UserID:        user2ID,
			PaymentMethod: "MOBILE_BANKING",
			Status:        "PENDING",
			Amount:        159.00,
			PaidAt:        now.Add(-4 * time.Hour),
			CreatedAt:     now.Add(-4 * time.Hour),
			UpdatedAt:     now.Add(-4 * time.Hour),
		},
	}

	receipts := []models.Receipt{
		{
			ID:            receipt1ID,
			ReceiptNo:     fmt.Sprintf("RCPT-%s-0001", now.Format("20060102")),
			UserID:        user1ID,
			BookingID:     booking1ID,
			PaymentID:     payment1ID,
			PaymentMethod: "CARD",
			Status:        "PAID",
			Amount:        129.00,
			PaidAt:        now.Add(-6 * time.Hour),
			CreatedAt:     now.Add(-6 * time.Hour),
			UpdatedAt:     now.Add(-5 * time.Hour),
		},
	}

	reviews := []models.Review{
		{
			ID:        review1ID,
			UserID:    user1ID,
			HotelID:   hotel1ID,
			Review:    "Clean rooms, friendly staff. Great value for money.",
			Rating:    8,
			CreatedAt: now.Add(-30 * time.Hour),
			UpdatedAt: now.Add(-29 * time.Hour),
		},
		{
			ID:        review2ID,
			UserID:    user2ID,
			HotelID:   hotel2ID,
			Review:    "Nice location near the river. Quiet and comfortable.",
			Rating:    7,
			CreatedAt: now.Add(-20 * time.Hour),
			UpdatedAt: now.Add(-19 * time.Hour),
		},
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := createAll(tx, images); err != nil {
			return err
		}
		if err := createAll(tx, admins); err != nil {
			return err
		}
		if err := createAll(tx, users); err != nil {
			return err
		}
		if err := createAll(tx, hotels); err != nil {
			return err
		}
		if err := createAll(tx, rooms); err != nil {
			return err
		}
		if err := createAll(tx, bookings); err != nil {
			return err
		}
		if err := createAll(tx, payments); err != nil {
			return err
		}
		if err := createAll(tx, receipts); err != nil {
			return err
		}
		if err := createAll(tx, reviews); err != nil {
			return err
		}
		return nil
	})
}

func createAll[T any](db *gorm.DB, rows []T) error {
	if len(rows) == 0 {
		return nil
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&rows).Error
}
