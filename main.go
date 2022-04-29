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

	// deprecated
	repo, err := database.NewRepo(env)
	if err != nil {
		log.Fatalln(err)
	}
	deprecatedCtr := controller.New(repo)

	repository, err := database.NewRepository(env)
	if err != nil {
		log.Fatal(err)
	}
	svc := appservice.NewService(repository)
	ctr := controller.NewController(svc)
	server := server.NewServer(env, ctr, deprecatedCtr)
	server.ApiV1Route()
	server.Run()
}
