package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBSource struct {
	Conn    *sql.DB
	Queries *Queries
}

func NewPostgres(dbdriver, dbsource string) (database *DBSource, err error) {
	sqlconn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		return
	}

	database.Conn = sqlconn
	database.Queries = New(sqlconn)
	return
}
