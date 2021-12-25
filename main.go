package main

import (
	"log"
	"try-bank/config"
	"try-bank/controller"
	"try-bank/database"
	"try-bank/server"
)

const (
	envPath = "."
)

func main() {
	env, err := config.NewEnv(envPath)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.NewRepository(env)
	if err != nil {
		log.Fatalln(err)
	}
	ctr := controller.NewController(db)

	server := server.NewServer(env, ctr)
	server.ApiRoute()
	server.Run()
}
