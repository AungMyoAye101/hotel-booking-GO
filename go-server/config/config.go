package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	SERVER   Server
	DATABASE Database
	AUTH     Auth
}
type Server struct {
	HOST string
	PORT string
}

type Database struct {
	URL      string
	Disabled bool
}

type Auth struct {
	ACCESS_SECRET  string
	REFRESH_SECRET string
}

func New() (*Config, error) {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.SetDefault("PORT", 8000)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	cfg := &Config{
		SERVER: Server{
			HOST: viper.GetString("HOST"),
			PORT: viper.GetString("PORT"),
		},
		DATABASE: Database{
			URL: viper.GetString("DATABASE_URL"),
		},
		AUTH: Auth{
			ACCESS_SECRET:  viper.GetString("ACCESS_TOKEN_SECRET"),
			REFRESH_SECRET: viper.GetString("REFRESH_TOKEN_SECRET"),
		},
	}
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
