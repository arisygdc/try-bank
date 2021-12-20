package main

import (
	"log"
	"try-bank/config"
	"try-bank/server"
)

const (
	envPath = "."
)

func main() {
	env, err := config.NewEnv(envPath)
	if err != nil {
		log.Println(err)
	}
	server := server.NewServer(env)
	server.WebRouteCustomConfig()
	server.WebRoute()
	server.ApiRoute()
	server.Run()
}
