package controller

import "try-bank/config"

type Controller struct {
	env config.Environment
}

func NewController(env config.Environment) (controller Controller) {
	controller = Controller{
		env: env,
	}
	return
}
