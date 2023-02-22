package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Berhtz/go-auth/dbconn"
	"github.com/Berhtz/go-auth/models"
	"github.com/Berhtz/go-auth/pkg/hash"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	models.DecodeRequest(&creds, r)
	db := dbconn.DbConnect()

	// check if user exists in database
	_, err := models.FindUserByUsername(creds.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			id := uuid.New()
			hashedPass, _ := hash.HashPassword(creds.Password)
			db.Create(&models.User{Id: id, Username: creds.Username,
				Password: hashedPass,
				Email:    creds.Email})
			log.Println("user created")
			return
		}
		log.Println("unexpected error while checking user in database")
		return
	} else {
		log.Println("User is already exist")
		return
	}

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	var user models.User

	models.DecodeRequest(&creds, r)

	// checking if user exists in db
	user, gormErr := models.FindUserByUsername(creds.Username)
	if gormErr == gorm.ErrRecordNotFound {
		fmt.Println("User not found")
		return
	} else if gormErr != nil {
		fmt.Println(gormErr)
		return
	}

	// Comparing given password with password in db
	fmt.Println("userPass:", user.Password, "creds pass:", creds.Password)
	err := hash.ComparePasswords(user.Password, creds.Password)
	if err == bcrypt.ErrHashTooShort {
		fmt.Println("Error while comparing passwords in SignIn func: ", err)
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User found, passwords matched, SignIn can be continued")
}
