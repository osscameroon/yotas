package auth

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{"templates/home.html"}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	redirectUri := env.BaseUri + "/github/login"
	scope := "read:user, user:email"

	authorizeUrl := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s",
		env.GithubClientId, redirectUri, scope,
	)

	http.Redirect(w, r, authorizeUrl, http.StatusTemporaryRedirect)
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	errMsg := r.FormValue("error")
	errDescription := r.FormValue("error_description")

	if errMsg != "" {
		fmt.Fprintf(w, errMsg, errDescription)
		return
	}

	// Get oauth code sent by github
	code := r.FormValue("code")

	// Get access token
	accessToken, err := NewGithubAccessToken(code)

	if err != nil || *accessToken == (GithubAccessToken{}) {
		fmt.Fprintf(w, "Unable to get access Token")
		return
	}

	// Retrieve user informations
	githubUser, err := NewGithubUser(accessToken)

	if err != nil {
		fmt.Fprintf(w, "Unable to get user Infos")
		return
	}

	// Retrieve user primary email
	githubEmail, err := NewGithubEmail(accessToken)

	if err != nil {
		fmt.Fprintf(w, "Unable to get user Email")
		return
	}

	// Check if a user with this email exists
	user := AppUser{
		Name:      githubUser.Name,
		Email:     githubEmail.Email,
		GithubId:  githubUser.Id,
		AvatarUrl: githubUser.AvatarUrl,
		LastLogin: time.Now(),
	}

	if !user.Exists() {
		db.Create(&user)
	}

	data := GithubProfilePageData{
		User:  *githubUser,
		Email: *githubEmail,
	}

	tmpFiles := []string{"templates/profile.html"}
	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	tmpl.Execute(w, data)
}
