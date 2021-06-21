package auth

import "github.com/gorilla/mux"

func AuthRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/github/login", GithubCallbackHandler).Methods("GET", "POST")

	return r
}
