package controller

import (
	"try-bank/database"
)

type Controller struct {
	Repo database.IRepo
}

func NewController(repository database.IRepo) (controller Controller) {
	controller = Controller{
		Repo: repository,
	}
	return
}
