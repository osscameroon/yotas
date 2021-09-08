package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func githubCallbackHandler(oauthConf oauth2.Config) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var call Callback
		err := c.BindJSON(&call)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 400,
				"reason": "Malformed request",
			})
			log.Println(err.Error())
		}

		token, err := oauthConf.Exchange(oauth2.NoContext, call.Code)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 400,
				"reason": "Invalid code",
			})
			log.Println(err.Error())
		}
		session := sessions.Default(c)
		session.Set("token", token.AccessToken)
		err = session.Save()
		// We return the user
		getUserFromToken(c, token.AccessToken, oauthConf)
	}

	return gin.HandlerFunc(fn)
}

func authHandler(oauthConf oauth2.Config) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// will always return the url
		Url := oauthConf.AuthCodeURL("hoge", oauth2.AccessTypeOnline)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"url":    Url,
		})
	}

	return gin.HandlerFunc(fn)
}


func githubOauthClient(c *gin.Context, token string, oauthConf oauth2.Config, fromCache bool) (*github.User, error) {
	var oauthClient *http.Client
	if fromCache {
		session := sessions.Default(c)
		tokenCached := session.Get("token")
		oauthClient = oauthConf.Client(oauth2.NoContext, &oauth2.Token{AccessToken: tokenCached.(string)})
	}else{
		oauthClient = oauthConf.Client(oauth2.NoContext, &oauth2.Token{AccessToken: token})
	}

	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(c, "")

	return user, err
}

func createUserHandler (user *github.User, token string){
	var err = CreateUser(*user, token)
	if err != nil {
		log.Println(err.Error())
	}
}

func getUserFromToken(c *gin.Context, token string, oauthConf oauth2.Config){

	var user *github.User
	var err error

	// we try to get from database
	githubUser, err := GetUserByToken(token)

	if err != nil {
		// we get from cache
		user, err = githubOauthClient(c, token, oauthConf, true)
		if err != nil {
			// we don't get from cache
			user, err = githubOauthClient(c, token, oauthConf, false)
			if err != nil {
				log.Println(err.Error())
			}else{
				createUserHandler(user, token)
			}
		}else{
			createUserHandler(user, token)
		}
	}else{
		user.Name = &githubUser.Name
		user.Email = &githubUser.Name
		user.NodeID = &githubUser.GithubId
		user.AvatarURL = &githubUser.AvatarUrl
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"user": gin.H{
			"id": 1,
			"github_profile": gin.H{
				"name": user.Name,
				"email": user.Email,
				"github_id": user.NodeID,
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
