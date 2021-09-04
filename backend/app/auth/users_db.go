package auth

import (
	"github.com/google/go-github/github"
	"github.com/osscameroon/yotas/db"
	"time"
)

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
