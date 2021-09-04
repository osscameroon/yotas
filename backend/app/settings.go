package app

import "os"

type env struct {
	HttpPort       string
	BaseUri        string
	PgHost         string
	PgDbName       string
	PgUser         string
	PgPassword     string
	PgPort         string
	GithubClientId string
	GithubClientSecret string
}

var Env *env

//InitEnv load env var and store it in Env
func InitEnv() {
	Env = &env{
		HttpPort:   os.Getenv("PORT"),
		BaseUri:    os.Getenv("BASE_URI"),
		PgHost:     os.Getenv("PG_HOST"),
		PgDbName:   os.Getenv("PG_DBNAME"),
		PgUser:     os.Getenv("PG_USER"),
		PgPassword: os.Getenv("PG_PASSWORD"),
		PgPort:     os.Getenv("PG_PORT"),
		GithubClientId:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret:     os.Getenv("GITHUB_CLIENT_SECRET"),
	}
}
