package auth

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/osscameroon/yotas/app"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
	"log"
	"os"
)

func AuthRouter() {
	router := app.GetApiRouter()

	// We get our secret
	secretKey := app.Env.SecretKey

	// auth configurations
	oauthConf := &oauth2.Config{
		ClientID:     app.Env.GithubClientId,
		ClientSecret: app.Env.GithubClientSecret,
		Scopes:       []string{"user"},
		Endpoint:     githuboauth.Endpoint,
	}

	if secretKey == "" {
		sk, err := generateSecretKey()
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				log.Println(err.Error())
			}
		}
		secretKey = sk
	}

	router.Use(sessions.Sessions("yotas", sessions.NewCookieStore([]byte(secretKey))))

	router.GET("/auth", authHandler(*oauthConf))
	router.POST("/auth", githubCallbackHandler(*oauthConf))
	router.GET("/users/:userID", getUserByIdHandler)
}
