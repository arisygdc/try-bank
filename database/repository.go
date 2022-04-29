package database

import (
	"context"
	"database/sql"
	"errors"
	"try-bank/config"
	"try-bank/database/postgresql"

	_ "github.com/lib/pq"
)

const PostgreDriver = "postgres"

var ErrNotFoundDriver = errors.New("database driver not found")

// implementing new repository
type IRepository interface {
	Query() *postgresql.Queries
	QueryTx(ctx context.Context, statement func(*postgresql.Queries) error) error
}

type Repository struct {
	sqlConn *sql.DB
	query   *postgresql.Queries
}

func NewRepository(env config.Environment) (IRepository, error) {
	if env.DBDriver == PostgreDriver {
		createSqlconn, err := sql.Open(env.DBDriver, env.DBSource)
		if err != nil {
			return Repository{}, err
		}

		createQuerier := postgresql.New(createSqlconn)
		return Repository{
			sqlConn: createSqlconn,
			query:   createQuerier,
		}, nil
	}
	return Repository{}, ErrNotFoundDriver
}

func (repos Repository) Query() *postgresql.Queries {
	return repos.query
}

func (repos Repository) QueryTx(ctx context.Context, statement func(*postgresql.Queries) error) error {
	tx, err := repos.sqlConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	QueryTx := repos.query.WithTx(tx)
	errorTx := statement(QueryTx)
	if errorTx != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}
		return errorTx
	}
	tx.Commit()
	return nil
}
