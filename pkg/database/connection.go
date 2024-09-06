package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"ocserv/pkg/config"
	"os"
)

var DB *gorm.DB

func RootConnect() *gorm.DB {
	cfg := config.GetDb()
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Connect() {
	cfg := config.GetDb()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func Connection() *gorm.DB {
	debug := os.Getenv("DEBUG") == "true"
	if debug {
		return DB.Debug()
	}
	return DB
}
