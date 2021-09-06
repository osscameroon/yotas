package auth

import (
	"github.com/google/go-github/github"
	"github.com/osscameroon/yotas/db"
	"time"
)

// CreateUser create a user from it's github user object and its token
func CreateUser(user github.User, token string) error {
	return db.Session.Create(&Users{
		Model			: db.Model{CreatedAt: time.Now().UTC()},
		Name 			: *user.Name,
		Email 			: *user.Email,
		GithubId 		: *user.NodeID,
		GithubToken 	: token,
		AvatarUrl 		: *user.AvatarURL,
		Active			: false,
	}).Error
}

//GetUserByID Retrieve a user from it's githubID
func GetUserByID(userID uint) (*Users, error) {
	var user Users
	result := db.Session.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

//GetUserByToken Retrieve a user from it's githubID
func GetUserByToken(token string) (*Users, error) {
	var user Users
	result := db.Session.Where("github_token = ?", token).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
