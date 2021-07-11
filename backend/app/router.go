package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var apiRouter *gin.RouterGroup

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
