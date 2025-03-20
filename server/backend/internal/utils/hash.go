package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(hashPassword, password string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}