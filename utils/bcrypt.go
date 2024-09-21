package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Bcrypt compare password with hash
func BcryptCompare(password string, hash string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false
		}

		fmt.Println("ERROR: " + err.Error())
	}
	return true
}
