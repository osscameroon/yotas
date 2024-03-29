package app

//GetOrganisationPicture Get a Pictures linked to an Organisations
func GetOrganisationPicture(pictureID, organisationID uint) (*Pictures, error) {
	var picture Pictures
	result := Session.Where("id = ? AND organisation_id = ?", pictureID, organisationID).First(&picture)
	if result.Error != nil {
		return nil, result.Error
	}

	return &picture, nil
}
