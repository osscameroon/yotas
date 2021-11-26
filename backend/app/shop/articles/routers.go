package articles

import (
	"github.com/osscameroon/yotas/app"
)

func ArticleRouter() {
	router := app.GetApiRouter()

	router.GET("/articles", GetOrganisationArticlesHandler)
	router.POST("/articles", app.IsAuthorized(CreateArticleHandler))
	router.GET("/articles/:articleID", GetArticleHandler)
	router.PUT("/articles/:articleID", app.IsAuthorized(UpdateArticleHandler))
	router.DELETE("/articles/:articleID", app.IsAuthorized(DeleteArticleHandler))
}
