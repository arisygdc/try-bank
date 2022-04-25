package controller

import (
	"try-bank/database"
)

type DeprecatedController struct {
	Repo database.IRepo
}

func NewController(repository database.IRepo) (controller DeprecatedController) {
	controller = DeprecatedController{
		Repo: repository,
	}
	return
}
