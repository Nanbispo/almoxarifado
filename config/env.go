package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("erro ao carregar arquivo .env: %w", err)
	}

	cfg := &Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	missing := make([]string, 0)
	for name, value := range map[string]string{
		"DB_HOST": cfg.DB_HOST,
		"DB_PORT": cfg.DB_PORT,
		"DB_USER": cfg.DB_USER,
		"DB_NAME": cfg.DB_NAME,
	} {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("variáveis de banco obrigatórias ausentes: %s", strings.Join(missing, ", "))
	}

	return cfg, nil
}
