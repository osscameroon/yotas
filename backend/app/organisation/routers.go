package organisation

import (
	"github.com/osscameroon/yotas/app"
)

func OrganisationRouter() {
	router := app.GetApiRouter()

	router.GET("/organisations", OrganisationsAllHandler)
	router.POST("/organisations", OrganisationsAllHandler)
	router.GET("/organisations/{organisation_id}", OrganisationsHandler)
	router.PATCH("/organisations/{organisation_id}", OrganisationsHandler)

	router.GET("/articles", GetOrganisationArticlesHandler)
	router.POST("/articles", CreateArticleHandler)
	router.GET("/articles/:articleID", GetArticleHandler)
	router.PUT("/articles/:articleID", UpdateArticleHandler)
	router.DELETE("/articles/:articleID", DeleteArticleHandler)

	router.GET("/orders", GetUserOrdersHandler)
	router.GET("/orders/organisation", GetOrganisationOrdersHandler)
	router.GET("/orders/:orderID", GetOrderHandler)
	router.POST("/orders", CreateOrderHandler)
	router.POST("/orders/:orderID/process", ProcessOrderHandler)
}
