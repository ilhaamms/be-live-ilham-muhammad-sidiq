package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	PortService string `envconfig:"PORT_SERVICE" required:"true"`
}

func NewAppConfig() (*AppConfig, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Error loading .env file " + err.Error())
	}

	var config AppConfig

	err = envconfig.Process("", &config)
	if err != nil {
		return nil, errors.New("Error loading environment variable " + err.Error())
	}

	return &config, nil
}
