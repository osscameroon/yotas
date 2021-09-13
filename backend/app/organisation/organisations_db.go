package organisation

import (
	"github.com/osscameroon/yotas/app"
)

//GetOrganisationPicture Get a Pictures linked to an Organisations
func GetOrganisationPicture(pictureID, organisationID uint) (*app.Pictures, error) {
	var picture app.Pictures
	result := app.Session.Where("id = ? AND organisation_id = ?", pictureID, organisationID).First(&picture)
	if result.Error != nil {
		return nil, result.Error
	}

	return &picture, nil
}
