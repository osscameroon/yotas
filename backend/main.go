package main

import (
	"log"
	"net/http"

	"github.com/osscameroon/yotas/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/osscameroon/yotas/app"
)

func main() {
	_ = godotenv.Load()

	app.InitEnv()

	db.Init()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// ─── API ROUTER ────────────────────────────────────────────────────
	apiRouter := router.Group(app.Env.BaseUri)

	apiRouterV1 := apiRouter.Group("/v1")
	apiRouterV1.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Yotas")
	})

	log.Println("HTTP Server Started on port ", app.Env.HttpPort)
	err := router.Run(":" + app.Env.HttpPort)
	log.Fatal(err)
}
