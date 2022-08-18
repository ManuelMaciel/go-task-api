package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost port=5432 user=postgres dbname=gorm password=i904010 sslmode=disable"

var DB *gorm.DB

func DBConnection() {
	// connect to database
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		log.Println("Connected to database")
	}
}
