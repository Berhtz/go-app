package main

import (
	"log"
	"net/http"

	"github.com/Berhtz/go-auth/auth/handlers"
	"github.com/Berhtz/go-auth/auth/models"
	"github.com/Berhtz/go-auth/dbconn"
	"github.com/joho/godotenv"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in auth/cmd/main.go")
	}

	//migrations
	db := dbconn.DbConnect()
	dbHasTable := db.Migrator().HasTable(&models.User{})
	if dbHasTable == false {
		db.Migrator().CreateTable(&models.User{})
	}

	http.HandleFunc("/api/signup", handlers.SignUp)
	http.HandleFunc("/api/signin", handlers.SignIn)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
