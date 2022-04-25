package controller

import (
	"try-bank/database"
)

// deprecated controller
type DeprecatedController struct {
	Repo database.IRepo
}

func New(repository database.IRepo) (controller DeprecatedController) {
	controller = DeprecatedController{
		Repo: repository,
	}
	return
}
