package testutils

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/config"
	"ocserv/pkg/database"
	tokenGenerator "ocserv/pkg/token"
	"os"
	"time"
)

func LoadTestEnv() {
	config.LoadEnv()
	err := os.Setenv("POSTGRES_NAME", "test")
	if err != nil {
		log.Fatal(err)
	}
}

func drop(db *gorm.DB, dbName string) {
	dropDBSQl := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)
	fmt.Println(dropDBSQl)
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
	LoadTestEnv()
	config.Set()
	database.Connect()
	return database.Connection()
}

func CreateTestAdminUser() *models.User {
	db := database.Connection()
	user := models.User{
		Username: "test-admin",
		Password: "test-admin-password",
		IsAdmin:  true,
	}
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	return &user
}

func CreateTestAdminToken(user *models.User) string {
	db := database.Connection()
	expireAt := time.Now().Add(time.Hour).Unix()
	key := tokenGenerator.Create(user.ID, expireAt)

	token := models.Token{
		UserID:   user.ID,
		ExpireAt: expireAt,
		Key:      key,
	}
	err := db.Create(&token).Error
	if err != nil {
		log.Fatal(err)
	}
	return token.Key
}
