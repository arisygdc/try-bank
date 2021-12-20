package database

import (
	"errors"
	"try-bank/config"
	"try-bank/database/postgres"
)

var (
	ErrNotFoundDriver = errors.New("database driver not found")
)

func NewSQL(env config.Environment) (dbsource *postgres.DB, err error) {
	if env.DBDriver == "postgres" {
		dbsource, err = postgres.NewPostgres(env.DBDriver, env.DBSource)
	} else {
		err = ErrNotFoundDriver
	}
	return
}
