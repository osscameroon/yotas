package auth

import (
	"github.com/google/go-github/github"
	"github.com/osscameroon/yotas/app"
	"time"
)

// CreateUser create a user from it's github user object and its token
func CreateUser(user github.User, githubToken string) error {
	return app.Session.Create(&app.Users{
		Model:       app.Model{CreatedAt: time.Now().UTC()},
		Name:        *user.Name,
		Email:       *user.Email,
		GithubId:    *user.NodeID,
		GithubToken: githubToken,
		AvatarUrl:   *user.AvatarURL,
	}).Error
}

//GetUserByID Retrieve a user from it's githubID
func GetUserByID(userID uint) (*app.Users, error) {
	var user app.Users
	result := app.Session.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

//GetUserByGithubToken Retrieve a user from it's githubID
func GetUserByGithubToken(token string) (*app.Users, error) {
	var user app.Users
	result := app.Session.Where("github_token = ?", token).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
