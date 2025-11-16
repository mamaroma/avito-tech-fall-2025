package config

import (
	"errors"
	"os"
)

type Config struct {
	DBConn string
}

func Load() (*Config, error) {
	conn := os.Getenv("DB_CONN")
	if conn == "" {
		return nil, errors.New("DB_CONN not set")
	}

	return &Config{
		DBConn: conn,
	}, nil
}
