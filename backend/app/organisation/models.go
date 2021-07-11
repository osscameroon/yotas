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

type Wallets struct {
	gorm.Model
	WalletId       string
	UserId         int64
	OrganisationId int64
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
	ArticleId int64
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
