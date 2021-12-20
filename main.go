package main

import (
	"log"
	"try-bank/config"
	"try-bank/server"
)

func main() {
	env, err := config.NewEnv(".")
	if err != nil {
		log.Println(err)
	}
	server := server.NewServer(env)
	server.WebRouteCustomConfig()
	server.WebRoute()
	server.ApiRoute()
	server.Run()
}
