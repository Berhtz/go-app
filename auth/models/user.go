package models

import (
	"errors"

	"github.com/Berhtz/go-auth/dbconn"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID
	Username string
	Password string
	Email    string
}

// rewrite to gofrs/uuid
func FindUserByUsername(username string) (User, error) {
	var user User
	db := dbconn.DbConnect()
	result := db.Where("username = ?", username).First(&user)
	//Find(&user, "username = ?", username)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errors.New("User not found")
		}
		return user, result.Error
	}
	return user, nil
}
