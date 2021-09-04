package auth

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/osscameroon/yotas/app"
	"os"
)

func AuthRouter() {
	router := app.GetApiRouter()

	if secretKey == "" {
		sk, err := generateSecretKey()
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		secretKey = sk
	}

	router.Use(sessions.Sessions("yotas", sessions.NewCookieStore([]byte(secretKey))))

	router.GET("/auth", authHandler)
	router.POST("/auth", githubCallbackHandler)
}
