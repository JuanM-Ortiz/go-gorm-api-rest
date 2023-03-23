package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const url = "host=localhost user=postgres password=fjoan57zoaik92k dbname=goapirest port=5432"

var DB *gorm.DB

func DbConnect() {
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}
