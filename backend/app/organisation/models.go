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
	OrganisationId uint
	ArticleId      uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type OrganisationsUsers struct {
	gorm.Model
	OrganisationId uint
	UserId         uint
	Active         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Wallets struct {
	gorm.Model
	WalletId       string
	UserId         uint
	OrganisationId uint
	Balance        int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Operations struct {
	gorm.Model
	Amount        int64
	WalletId      string
	OperationType string
	Approved      bool
	OperationHash string
	CreatedAt     time.Time
}

type Orders struct {
	gorm.Model
	WalletId  string
	ArticleId uint
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Articles struct {
	gorm.Model
	Name        string
	Description string
	Quantity    int64
	Price       int64
	Metadata    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Pictures struct {
	gorm.Model
	OrganisationId uint
	AltText        string
	Original       string
	Small          string
	Medium         string
	Large          string
}

type ArticlesPictures struct {
	gorm.Model
	PictureId uint
	ArticleId uint
}
