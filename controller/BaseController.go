package controller

import (
	"try-bank/database/postgres"
)

type Controller struct {
	DBSource *postgres.DB
}

func NewController(dbsource *postgres.DB) (controller Controller) {
	controller = Controller{
		DBSource: dbsource,
	}
	return
}
