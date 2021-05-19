package wallet

import (
	"gorm.io/gorm"
	"time"
)

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
