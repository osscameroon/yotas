package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)
// auth configurations
var oauthConf = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Scopes:       []string{"user"},
	Endpoint:     githuboauth.Endpoint,
}
var secretKey = os.Getenv("SECRET_KEY_BASE")

type UsersPresenter struct {
	Users
}

type Callback struct {
	Code string `json:"code"`
	State string `json:"state"`
}

func githubCallbackHandler(c *gin.Context){
	var call Callback
	err := c.BindJSON(&call)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 400,
			"reason": "Malformed request",
		})
		return
	}

	token, err1 := oauthConf.Exchange(oauth2.NoContext, call.Code)
	if err1 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 400,
			"reason": "Invalid code",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("token", token.AccessToken)
	err1 = session.Save()
	// We return the user
	getUserFromToken(c, token.AccessToken)
}

func authHandler(c *gin.Context) {
	// will always return the url
	Url := oauthConf.AuthCodeURL("hoge", oauth2.AccessTypeOnline)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"url":  Url,
	})
}

func getUserFromToken(c *gin.Context, token string){
	oauthClient := oauthConf.Client(oauth2.NoContext, &oauth2.Token{AccessToken: token})
	client := github.NewClient(oauthClient)

	user, _, err := client.Users.Get(c, "")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"reason": "User not found",
		})
		return
	}
	// the user creation
	err = CreateUser(*user, token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"user": gin.H{
			"id": 1,
			"github_profile": gin.H{
				"name": user.Name,
				"email": user.Email,
				"github_id": user.ID,
				"avatar_url": user.AvatarURL,
			},
			"active": true,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
		"token": token,
	})
}

func generateSecretKey() (string, error) {
	b := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, b)

	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}

	return strings.TrimRight(base64.StdEncoding.EncodeToString(b), "="), nil
}
