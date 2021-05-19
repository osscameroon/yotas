package shop

import (
	"gorm.io/gorm"
	"time"
)

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
