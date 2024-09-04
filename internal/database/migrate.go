package database

import (
	"gorm.io/gorm"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/config"
	"ocserv/pkg/database"
	"strings"
)

func migrateEnums(db *gorm.DB) {
	err := db.Exec(`CREATE TYPE service_type_enum AS ENUM ('FREE', 'MONTHLY', 'TOTALLY');`).Error
	if err != nil {
		if strings.Contains(err.Error(), "type \"service_type_enum\" already exists") {
			log.Println("Type \"service_type_enum\" Already Exists. Continue Migrating ...")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Type \"service_type_enum\" Created Successfully ...")
	}
}

func Migrate() {
	config.Set()
	database.Connect()
	db := database.Connection()
	migrateEnums(db)
	err := db.AutoMigrate(&models.Site{}, &models.User{}, &models.Token{}, &models.OcservUser{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Migration Completed Successfully ...")
}
