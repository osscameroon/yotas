package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/osscameroon/yotas/app"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// getUserByIdHandler will get the user from the given userID
func getUserByIdHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "User id must be an int",
		})
		return
	}

	user, err := GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": err.Error(),
		})
		return
	}
	// Not sure yet about this, but the github-id and the gthub-token
	// should be private right ?
	user.GithubId = ""
	user.GithubToken = ""

	result := app.UsersPresenter{Users: *user}

	c.JSON(http.StatusOK, result)
}

// githubCallbackHandler
func githubCallbackHandler(oauthConf oauth2.Config) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var call app.Callback
		err := c.BindJSON(&call)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 400,
				"reason": "Malformed request",
			})
			log.Println(err.Error())
			return
		}

		token, err := oauthConf.Exchange(context.TODO(), call.Code)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 400,
				"reason": "Invalid code",
			})
			log.Println(err.Error())
			return
		}

		session := sessions.Default(c)
		session.Set("token", token.AccessToken)
		err = session.Save()
		if err != nil {
			log.Println(err.Error())
			return
		}

		// We return the user
		getUserFromGithubToken(c, token.AccessToken, oauthConf)
	}

	return gin.HandlerFunc(fn)
}

// authHandler will return the url to the client with the githubClientID for him to logIN
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
		oauthClient = oauthConf.Client(context.TODO(), &oauth2.Token{AccessToken: tokenCached.(string)})
	} else {
		oauthClient = oauthConf.Client(context.TODO(), &oauth2.Token{AccessToken: token})
	}

	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(c, "")

	return user, err
}

// getUserFromGithubToken this function will get a user from github token
func getUserFromGithubToken(c *gin.Context, githubToken string, oauthConf oauth2.Config) {
	// we try to get from database
	githubUser, err := GetUserByGithubToken(githubToken)

	if err != nil {
		// we get from cache
		githubUser, err := githubOauthClient(c, githubToken, oauthConf, true)
		if err != nil {
			// we don't get from cache
			githubUser, err := githubOauthClient(c, githubToken, oauthConf, false)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{ "reason": err.Error()})
			} else {
				err := CreateUser(*githubUser, githubToken)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{ "reason": err.Error()})
				}
			}
		} else {
			err := CreateUser(*githubUser, githubToken)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{ "reason": err.Error()})
			}
		}
	}

	// we generate the JWT token that we will return to the user
	jwtValidToken, err := app.GenerateJWT(githubToken)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": "Failed to generate Jwt token: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"user": gin.H{
			"id": 1,
			"github_profile": gin.H{
				"name":       &githubUser.Name,
				"email":      &githubUser.Email,
				"github_id":  &githubUser.GithubId,
				"avatar_url": &githubUser.AvatarUrl,
			},
			"active":     true,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
		"token": jwtValidToken,
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


