package config

import (
	"fmt"

	"github.com/utamadeveloper/golang-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_HOST = "localhost"
var DB_PORT = 5432
var DB_NAME = "order_api"
var DB_PASS = "root"
var DB_USER = "postgres"
var TZ = "Asia/Makassar"
var Psql *gorm.DB

func PostgresConnect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		DB_HOST,
		DB_USER,
		DB_PASS,
		DB_NAME,
		DB_PORT,
		TZ,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Psql = db

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection success")
	}

	Psql.Debug().AutoMigrate(&models.User{}, &models.Order{})
}

func GetDb() *gorm.DB {
	return Psql
}
