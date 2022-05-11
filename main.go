package main

import (
	"log"
	appservice "try-bank/app_service"
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

	repository, err := database.NewRepository(env)
	if err != nil {
		log.Fatal(err)
	}

	service := appservice.NewService(repository)
	controller := controller.NewController(service)
	server := server.NewServer(env, controller)
	server.ApiV1Route()
	server.Run()
}
