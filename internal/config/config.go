package config

import (
	"github.com/spf13/viper"
)

type (
	DB struct {
		Host     string
		Port     int
		DBName   string
		Username string
		Password string
		SSLMode  string
	}

	SpaceX struct {
		ApiUrl string
	}

	Server struct {
		Port int
	}

	Config struct {
		DB         DB
		SpaceX     SpaceX
		HTTPServer Server
	}
)

func ReadConfig() (Config, error) {
	viper.AddConfigPath("/app")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
