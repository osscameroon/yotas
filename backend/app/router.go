package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var router *gin.Engine
var apiRouter *gin.RouterGroup
var jwtSigningKey = []byte("osscameroonisamazing")

func InitRouter() {
	router = gin.Default()
	apiRouter = router.Group(Env.BaseUri)
}

func GetRouter() *gin.Engine {
	return router
}

func GetApiRouter() *gin.RouterGroup {
	return apiRouter
}

// GenerateJWT will generate the JWT token
func GenerateJWT(customKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = customKey
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(jwtSigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// IsAuthorized This function will check the authorization from a given token in Bearer
func IsAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			token, err := jwt.Parse(c.GetHeader("Authorization"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return jwtSigningKey, nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"reason": err.Error(),
				})
			}
			if token.Valid {
				endpoint(c)
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"reason": "Not Authorized",
			})
		}
	}
	return gin.HandlerFunc(fn)
}
