package app

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var apiRouter *gin.RouterGroup

func InitRouter() {
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
	}))
	apiRouter = router.Group(Env.BaseUri)
}

func GetRouter() *gin.Engine {
	return router
}

func GetApiRouter() *gin.RouterGroup {
	return apiRouter
}
