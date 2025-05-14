package dbl

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"order-api/config"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf *config.Config) *DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		conf.DB.Host,
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.DBName,
		conf.DB.Port,
		conf.DB.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &DB{db}
}
