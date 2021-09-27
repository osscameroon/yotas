package articles

import (
	"github.com/osscameroon/yotas/app"
)

func ArticleRouter() {
	router := app.GetApiRouter()

	router.GET("/articles", GetOrganisationArticlesHandler)
	router.POST("/articles", CreateArticleHandler)
	router.GET("/articles/:articleID", GetArticleHandler)
	router.PUT("/articles/:articleID", UpdateArticleHandler)
	router.DELETE("/articles/:articleID", DeleteArticleHandler)
}
