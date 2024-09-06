package testutils

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"ocserv/pkg/config"
	"ocserv/pkg/database"
)

func drop(db *gorm.DB, dbName string) {
	dropDBSQl := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)
	log.Println(dropDBSQl)
	err := db.Exec(dropDBSQl).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("database %s dropped", dbName)
}

func create(db *gorm.DB, dbName string) {
	createDBSQL := fmt.Sprintf("CREATE DATABASE %s", dbName)
	err := db.Exec(createDBSQL).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("database %s created", dbName)
}

func DropAndCreateDB(dbName string) {
	db := database.RootConnect()
	drop(db, dbName)
	create(db, dbName)
	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *sql.DB) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
}

func GetTestDB() *gorm.DB {
	config.LoadTestEnv()
	config.Set()
	database.Connect()
	return database.Connection()
}
