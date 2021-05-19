package auth

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	Name      string
	Email     string
	GithubId  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}
