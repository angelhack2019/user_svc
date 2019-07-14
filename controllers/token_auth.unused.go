package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const (
	Secret            = "secret"
	MockPassword      = "password"
	ExpirationMinutes = 5
)

type CustomClaims struct {
	Username string `json:username`
	jwt.StandardClaims
}

func secret() []byte {
	return []byte(Secret)
}

func minutes(num int) int64 {
	return time.Now().Add(time.Duration(num) * time.Minute).Unix()
}

func authenticate(username string, password string) bool {
	// TODO: check database for password
	return MockPassword == password
}

func token(username string, expiration int64) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration,
		},
	})
}

func get_token_string(username string, password string) (string, int) {
	if !authenticate(username, password) {
		return "", http.StatusUnauthorized
	}

	expiration := minutes(ExpirationMinutes)
	tokenString, err := token(username, expiration).SignedString(secret())

	if err != nil {
		return "", http.StatusInternalServerError
	}

	return tokenString, http.StatusOK
}

func authenticate_token(tokenString string) int {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret(), nil
	})

	if !token.Valid {

	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized
		}
		return http.StatusBadRequest
	}

	return http.StatusOK
}

func refresh_token(){
	// WOULD HAVE BEEN DONE BUT CJ HAPPENED
}