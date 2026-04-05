// config/validate.go
package config

import (
	"errors"
	"net/url"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func validateConfig(cfg *Config) error {

	err := validate.Struct(struct {
		DATABASE_URL   string `validate:"required"`
		PORT           string `validate:"required"`
		ACCESS_SECRET  string `validate:"required"`
		REFRESH_SECRET string `validate:"required"`
	}{

		DATABASE_URL:   cfg.DATABASE.URL,
		PORT:           cfg.SERVER.PORT,
		ACCESS_SECRET:  cfg.AUTH.ACCESS_SECRET,
		REFRESH_SECRET: cfg.AUTH.REFRESH_SECRET,
	})

	if err != nil {
		return errors.New("invalid configuration: " + err.Error())
	}

	parsed, parseErr := url.Parse(cfg.DATABASE.URL)
	if parseErr != nil || parsed.Scheme == "" || parsed.Host == "" {
		return errors.New("invalid configuration: DATABASE_URL must be a valid URL (e.g. postgresql://user:pass@host:5432/dbname)")
	}

	return nil
}
