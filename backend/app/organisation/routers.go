package organisation

import (
	"github.com/osscameroon/yotas/app"
)

func OrganisationRouter() {
	router := app.GetApiRouter()

	router.GET("/organisations", OrganisationsAllHandler)
	router.POST("/organisations", app.IsAuthorized(CreateOrganisationHandler))
	router.GET("/organisations/:organisationID", OrganisationsHandler)
	// router.PATCH("/organisations/:organisationID", UpdateOrganisationsHandler)
}
