package dbconn

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbConnect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("Postgres_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
