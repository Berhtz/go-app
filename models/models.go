package models

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
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

type Credentials struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username" validate:"required, min=4, max=16"`
	Password string    `json:"password" validate:"required, min=8, max=32"`
	Email    string    `json:"email" validate:"required, email"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func DecodeRequest(creds *Credentials, r *http.Request) error {
	return json.NewDecoder(r.Body).Decode(&creds)
}
