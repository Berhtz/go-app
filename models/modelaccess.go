package models

import (
	"log"

	"github.com/Berhtz/go-auth/dbconn"
)

func FindUserByUsername(username string) (User, error) {
	var user User
	db := dbconn.DbConnect()
	result := db.Where("username = ?", username).First(&user)
	log.Println("FindUserByUsername result:", result, result.Error, user)
	return user, result.Error
}
