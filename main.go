package main

import (
	"log"
	"try-bank/api"
	"try-bank/config"
)

func main() {
	env, err := config.NewEnv(".")
	if err != nil {
		log.Println(err)
	}
	server := api.NewServer(env)
	routes := server.RouteGroup("")
	routes.GET("/")
}
