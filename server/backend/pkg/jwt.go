package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("mysecretkey")

func GenerateJWT(username string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":username,
		"exp": time.Now().Add(time.Hour*24).Unix(),
	})
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		return claims, nil
	}
	return nil, err
}

