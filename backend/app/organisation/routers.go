package organisation

import "github.com/gorilla/mux"

func OrganisationRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/organisations", OrganisationsAllHandler).Methods("GET", "POST")
	r.HandleFunc("/organisations/{organisation_id}", OrganisationsHandler).Methods("GET", "PATCH")

	return r
}
