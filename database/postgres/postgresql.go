package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn    *sql.DB
	Queries *Queries
}

func NewPostgres(dbdriver, dbsource string) (database *DB, err error) {
	sqlconn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		return
	}

	database.Conn = sqlconn
	database.Queries = New(sqlconn)
	return
}
