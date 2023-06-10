package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenToken(user User, key string) string {
	jwt_token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		"roles":    user.Role,
		"username": user.Username,
	}
	// Sign and get the complete encoded token as a string
	token, err := jwt_token.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err, key)
	}
	return token
}
