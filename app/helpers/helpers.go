package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) []byte {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err == nil {
		return hash
	}

	panic("Cannot generate the password")
}

func GenerateToken() string {
	h := sha256.New()
	h.Write([]byte(RandStringBytes(64)))
	return hex.EncodeToString(h.Sum(nil))
}

func PasswordIsValid(password string, encryptedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(encryptedPassword, []byte(password))
	return err == nil
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ClearStruct(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}
