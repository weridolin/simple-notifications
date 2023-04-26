package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, id uint) {
	var user database.User
	if id != 0 {
		db := database.GetDB()
		db.First(&user, id)
	}
	c.Set("userID", id)
	c.Set("user", user)
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(database.JWT_KEY))
			return b, nil
		})
		if err != nil {
			fmt.Println(err)
			common.HttpResponse(c, http.StatusUnauthorized, -1, err.Error(), nil)
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			id := uint(claims["id"].(float64))
			UpdateContextUserModel(c, id)
		}
	}
}
