package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	dsn := "host=postgres user=postgres password=xlr22856 dbname=practice port=5432 "
	Instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
