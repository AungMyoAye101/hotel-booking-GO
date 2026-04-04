// config/validate.go
package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func validateConfig(cfg *Config) error {
	err := validate.Struct(struct {
		DATABASE_URL string `validate:"required,url"`
		PORT         string `validate:"required"`
		// JWTSecret   string `validate:"required,min=10"`
	}{

		DATABASE_URL: cfg.DATABASE_URL,
		PORT:         cfg.PORT,
		// JWTSecret:   cfg.JWTSecret,
	})

	if err != nil {
		return errors.New("invalid configuration: " + err.Error())
	}

	return nil
}
