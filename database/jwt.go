package database

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWT_KEY = "123456"

func GenToken(user User) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		"roles":    user.Role,
		"username": user.Username,
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(JWT_KEY))
	return token
}
