package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DB   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

type AuthConfig struct {
	Secret     string
	VerifyCode int
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	verifyCode, err := strconv.Atoi(os.Getenv("AUTH_VERIFY_CODE"))
	if err != nil {
		log.Fatal("Error loading AUTH_VERIFY_CODE")
	}

	authSecret := os.Getenv("AUTH_SECRET")
	if authSecret == "" {
		log.Fatal("AUTH_SECRET is empty")
	}

	return &Config{
		DB: DbConfig{
			Host:     os.Getenv("DB_HOST"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		Auth: AuthConfig{
			Secret:     authSecret,
			VerifyCode: verifyCode,
		},
	}
}
