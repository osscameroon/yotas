package auth

import (
	"github.com/osscameroon/yotas/db"
)

type Users struct {
	db.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	GithubId    string `json:"github_id"`
	GithubToken string `json:"github_token"`
	AvatarUrl   string `json:"avatar_url"`
}
