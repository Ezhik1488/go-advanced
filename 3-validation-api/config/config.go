package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Email EmailConfig
}

type EmailConfig struct {
	Login    string
	Password string
	Address  string
	Port     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}
	return &Config{
		Email: EmailConfig{
			Login:    os.Getenv("EMAIL_LOGIN"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Address:  os.Getenv("EMAIL_ADDRESS"),
			Port:     os.Getenv("EMAIL_PORT"),
		},
	}
}
