package config

import (
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	SERVER   Server
	DATABASE Database
}
type Server struct {
	HOST string
	PORT string
}

type Database struct {
	URL      string
	Disabled bool
}

func New() (*Config, error) {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.SetDefault("PORT", 8000)
	viper.SetDefault("SKIP_DB", false)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	// viper.GetBool reads strings like "true"/"false" from env, but when values
	// come from an .env file, they're read as strings; normalize here.
	skipDB := viper.GetBool("SKIP_DB")
	if raw := viper.GetString("SKIP_DB"); raw != "" {
		if parsed, err := strconv.ParseBool(raw); err == nil {
			skipDB = parsed
		}
	}

	cfg := &Config{
		SERVER: Server{
			HOST: viper.GetString("HOST"),
			PORT: viper.GetString("PORT"),
		},
		DATABASE: Database{
			URL:      viper.GetString("DATABASE_URL"),
			Disabled: skipDB,
		},
	}
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
