package pgrepo

import (
	"context"
	"database/sql"
	"try-bank/database/postgres"
)

type DB struct {
	conn    *sql.DB
	queries *postgres.Queries
}

func NewPostgres(dbdriver, dbsource string) (database *DB, err error) {
	sqlconn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		return
	}

	database = &DB{
		conn:    sqlconn,
		queries: postgres.New(sqlconn),
	}
	return
}

func (d DB) transaction(ctx context.Context, queryFunc func(*postgres.Queries) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queriesTx := d.queries.WithTx(tx)
	err = queryFunc(queriesTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}
