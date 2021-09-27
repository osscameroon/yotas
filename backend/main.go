package main

import (
	"github.com/osscameroon/yotas/app"
	"github.com/osscameroon/yotas/app/auth"
	"github.com/osscameroon/yotas/app/organisation"
	"github.com/osscameroon/yotas/app/shop/articles"
	"github.com/osscameroon/yotas/app/shop/orders"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	app.InitEnv()

	app.Init()

	// Init global router
	app.InitRouter()

	//Init other module router to bind with global router
	auth.AuthRouter()
	organisation.OrganisationRouter()
	orders.OrderRouter()
	articles.ArticleRouter()

	log.Println("HTTP Server Started on port ", app.Env.HttpPort)
	err := app.GetRouter().Run(":" + app.Env.HttpPort)
	if err != nil {
		log.Fatal(err)
	}
}
