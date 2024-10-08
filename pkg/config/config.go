package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	App App
	Db  Db
}

type App struct {
	Debug        bool
	Host         string
	Port         string
	SecretKey    string
	AllowOrigins []string
	Hook         bool
}

type Db struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

var conf Config

func LoadEnv() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envFile := filepath.Join(pwd, ".env")
	if _, err := os.Stat(envFile); !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func Set() {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable not set")
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := App{
		SecretKey:    secretKey,
		Host:         host,
		Port:         port,
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ", "),
	}

	debug, err := strconv.Atoi(os.Getenv("DEBUG"))
	if err != nil {
		app.Debug = true
	} else {
		app.Debug = debug == 1
	}

	hook, err := strconv.Atoi(os.Getenv("HOOK"))
	if err != nil {
		app.Hook = true
	} else {
		app.Hook = hook == 1
	}

	db := Db{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Name:     os.Getenv("POSTGRES_NAME"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	if ssl := os.Getenv("POSTGRES_SSL_MODE"); ssl == "require" {
		db.SSLMode = "require"
	} else {
		db.SSLMode = "disable"
	}

	if test := os.Getenv("TEST"); test == "true" {
		db.Name = "test"
	}

	conf = Config{
		App: app,
		Db:  db,
	}
}

func GetApp() *App {
	return &conf.App
}

func GetDb() *Db {
	return &conf.Db
}
