package handlers

import (
	"fmt"
	"net/http"

	"github.com/Berhtz/go-auth/auth/hash"
	"github.com/Berhtz/go-auth/auth/models"
	"github.com/Berhtz/go-auth/dbconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	var user models.User

	models.DecodeRequest(&creds, r)

	// Searching for given username in db
	db := dbconn.DbConnect()
	gormErr := db.Where("username = ?", creds.Username).First(&user).Error
	if gormErr == gorm.ErrRecordNotFound {
		fmt.Println("User not found")
		return
	} else if gormErr != nil {
		fmt.Println(gormErr)
		return
	}

	// Compating given password with password in db
	err := hash.ComparePasswords(user.Password, creds.Password)
	if err == bcrypt.ErrHashTooShort {
		fmt.Println("Check given hashed password in SignIn func: ", err)
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User found, passwords matched, SignIn can be continued")
}
