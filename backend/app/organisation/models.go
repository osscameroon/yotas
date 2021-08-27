package organisation

import (
	"github.com/osscameroon/yotas/db"
)

type Organisations struct {
	db.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	GithubId    string `json:"github_id"`
	AvatarUrl   string `json:"avatar_url"`
	WebSite     string `json:"web_site"`
	Description string `json:"description"`
}

type OrganisationsArticles struct {
	db.Model
	OrganisationId uint `json:"organisation_id"`
	ArticleId      uint `json:"article_id"`
}

type OrganisationsUsers struct {
	db.Model
	OrganisationId uint `json:"organisation_id"`
	UserId         uint `json:"user_id"`
	Active         bool `json:"active"`
}

type Wallets struct {
	db.Model
	WalletId       string `json:"wallet_id"`
	UserId         uint   `json:"user_id"`
	OrganisationId uint   `json:"organisation_id"`
	Balance        int64  `json:"balance"`
}

type Operations struct {
	db.Model
	Amount        int64  `json:"amount"`
	WalletId      string `json:"wallet_id"`
	OperationType string `json:"operation_type"`
	Approved      bool   `json:"approved"`
	OperationHash string `json:"operation_hash"`
}

type Orders struct {
	db.Model
	WalletId string `json:"wallet_id"`
	State    string `json:"state"`
	Decision string `json:"decision"`
}

type OrdersArticles struct {
	db.Model
	OrderID   uint `json:"order_id"`
	ArticleID uint `json:"article_id"`
	Quantity  int  `json:"quantity"`
}

type Articles struct {
	db.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
	Metadata    string `json:"metadata"`
}

type Pictures struct {
	db.Model
	OrganisationId uint   `json:"organisation_id"`
	AltText        string `json:"alt_text"`
	Original       string `json:"original"`
	Small          string `json:"small"`
	Medium         string `json:"medium"`
	Large          string `json:"large"`
}

type ArticlesPictures struct {
	db.Model
	PictureId uint `json:"picture_id"`
	ArticleId uint `json:"article_id"`
}
