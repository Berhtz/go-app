package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Generate a hash from the password using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPassword string, password string) error {
	// Compare the hashed password with the input password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
