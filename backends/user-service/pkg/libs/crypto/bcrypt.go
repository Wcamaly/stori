package crypto

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func validatePassword(hashPwd string, pwd []byte) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), pwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
