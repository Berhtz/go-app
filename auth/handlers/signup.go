package handlers

import (
	"fmt"
	"net/http"

	"github.com/Berhtz/go-auth/auth/hash"
	"github.com/Berhtz/go-auth/auth/models"
	"github.com/Berhtz/go-auth/dbconn"
	"github.com/google/uuid"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	models.DecodeRequest(&creds, r)
	db := dbconn.DbConnect()

	_, e := models.FindUserByUsername(creds.Username)
	if e != nil {
		id := uuid.New()
		hashedPass, _ := hash.HashPassword(creds.Password)
		db.Create(&models.User{Id: id, Username: creds.Username,
			Password: hashedPass,
			Email:    creds.Email})
		return
	} else {
		fmt.Println("User already exist")
		return
	}
}
