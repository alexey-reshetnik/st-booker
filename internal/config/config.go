package config

import (
	"os"

	"github.com/spf13/viper"
)

type (
	DB struct {
		Host     string
		Port     int
		DBName   string
		Username string
		Password string
		SSLMode  bool
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
	viper.SetConfigFile(os.Getenv("CONFIG_FILE"))

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
