package main

import (
	"github.com/osscameroon/yotas/app/auth"
	"github.com/osscameroon/yotas/app/organisation"
	"github.com/osscameroon/yotas/db"
	"log"

	"github.com/joho/godotenv"
	"github.com/osscameroon/yotas/app"
)

func main() {
	_ = godotenv.Load()

	app.InitEnv()

	db.Init()

	// Init global router
	app.InitRouter()

	//Init other module router to bind with global router
	auth.AuthRouter()
	organisation.OrganisationRouter()

	log.Println("HTTP Server Started on port ", app.Env.HttpPort)
	err := app.GetRouter().Run(":" + app.Env.HttpPort)
	if err != nil {
		log.Fatal(err)
	}
}
