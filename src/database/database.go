package database

import (
	"fmt"
	"log"

	"github.com/ItzRyz/squareville-api/src/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	host := lib.GetEnvVar("DB_HOST")
	port := lib.GetEnvVar("DB_PORT")
	username := lib.GetEnvVar("DB_USERNAME")
	password := lib.GetEnvVar("DB_PASSWORD")
	database := lib.GetEnvVar("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database")
	}
	fmt.Println("Connecting to database")
}
