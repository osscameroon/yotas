package db

import (
	"fmt"
	"log"

	"github.com/osscameroon/yotas/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Session is the instance of the database connection
var Session *gorm.DB

//Init start the connection to the database
func Init() {

	var err error

	params := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		app.Env.PgHost, app.Env.PgUser, app.Env.PgPassword, app.Env.PgDbName, app.Env.PgPort,
	)

	Session, err = gorm.Open(postgres.Open(params), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL server", err)
	}

	log.Println("Connected to PostgreSQL Database")
}
