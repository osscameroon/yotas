package app

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

//Session is the instance of the database connection
var Session *gorm.DB

//Init start the connection to the database
func Init() {
	var err error

	params := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Env.PgHost, Env.PgUser, Env.PgPassword, Env.PgDbName, Env.PgPort,
	)

	Session, err = gorm.Open(postgres.Open(params), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL server", err)
	}

	log.Println("Connected to PostgreSQL Database")
}
