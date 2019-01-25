package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"
)

func GenerateToken(password string) ([]byte, string) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err == nil {
		return hash, base64.StdEncoding.EncodeToString(hash)
	} else {
		panic("Cannot generate the password")		
	}
}