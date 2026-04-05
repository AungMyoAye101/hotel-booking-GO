package auth

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}
