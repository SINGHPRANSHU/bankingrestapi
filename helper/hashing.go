package helper

import (
	"golang.org/x/crypto/bcrypt"
)

// CreateHash Will create hash password
// It should never panic if plainText is given properly
func CreateHash(plainText string) (hashText string) {
	
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashText = string(passwordHashInBytes)
	return 
}

// CompareHash compares hash to plain text, if same it will no return error
// If not same, it will return error
func CompareHash(plainText string, hashText string) (err error) {
	plainTextInBytes := []byte(plainText)
	
	hashTextInBytes := []byte(hashText)
	err = bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
	return
}