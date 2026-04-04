package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string
	PORT         string
}

func LoadConfig() (*Config, error) {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.SetDefault("PORT", "8000")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	cfg := &Config{
		DATABASE_URL: viper.GetString("DATABASE_URL"),
		PORT:         viper.GetString("PORT"),
	}
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}
	log.Println("env loaded")
	return cfg, nil
}
