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
}
