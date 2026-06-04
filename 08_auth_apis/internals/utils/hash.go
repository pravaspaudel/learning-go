package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//returns nil when comparision is false
func ComparePassword(plainText string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
}
