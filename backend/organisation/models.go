package organisation

import (
	"gorm.io/gorm"
	"time"
)

type Organisations struct {
	gorm.Model
	Name        string
	Email       string
	GithubId    string
	AvatarUrl   string
	WebSite     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrganisationsArticles struct {
	gorm.Model
	OrganisationId int64
	ArticleId      int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type OrganisationsUsers struct {
	gorm.Model
	OrganisationId int64
	UserId         int64
	Active         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
