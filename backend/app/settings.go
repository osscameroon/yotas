package app

import "os"

type env struct {
	HttpPort string
	BaseUri  string
}

var Env *env

//InitEnv load env var and store it in Env
func InitEnv() {
	Env = &env{
		HttpPort: os.Getenv("PORT"),
		BaseUri:  os.Getenv("BASE_URI"),
	}
}
