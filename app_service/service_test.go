package appservice

import (
	"try-bank/config"
)

func getDBConf() config.Environment {
	return config.Environment{
		DBDriver: "postgres",
		DBSource: "postgresql://postgresTest:secret@localhost:5432/bank?sslmode=disable",
	}
}
