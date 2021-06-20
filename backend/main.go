package main

import (
	"github.com/osscameroon/yotas/db"
	"log"
	"net/http"

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
	apiRouter := router.Group("/api")

	apiRouterV1 := apiRouter.Group("/v1")
	apiRouterV1.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Yotas")
	})

	log.Println("HTTP Server Started on port ", app.Env.HttpPort)
	err := router.Run(":" + app.Env.HttpPort)
	log.Fatal(err)
}
