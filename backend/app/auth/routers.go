package auth

import (
	"github.com/osscameroon/yotas/app"
)

func authRouter() {
	router := app.GetApiRouter()

	router.POST("/login", LoginHandler)
	router.GET("/github/login", GithubCallbackHandler)
	router.POST("/github/login", GithubCallbackHandler)
}
