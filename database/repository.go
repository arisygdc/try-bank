package database

import (
	"context"
	"database/sql"
	"errors"
	"try-bank/config"
	"try-bank/database/postgres/pgrepo"
	"try-bank/database/postgresql"
	"try-bank/request"

	_ "github.com/lib/pq"
)

const PostgreDriver = "postgres"

var ErrNotFoundDriver = errors.New("database driver not found")

type IRepo interface {
	CreateLevel(ctx context.Context, req request.PermissionReq) error
	CreateUser(ctx context.Context, req request.PostUser, permission string) (int32, error)
	CreateCompany(ctx context.Context, req request.PostCompany) (int32, error)
	ActivateVA(ctx context.Context, req request.VirtualAccount) (int32, string, error)
	PaymentVA(ctx context.Context, req request.PaymentVA) error
	Transfer(ctx context.Context, req request.Transfer) error
}

// deprecated repository
func NewRepo(env config.Environment) (dbsource IRepo, err error) {
	if env.DBDriver == "postgres" {
		dbsource, err = pgrepo.NewPostgres(env.DBDriver, env.DBSource)
	} else {
		err = ErrNotFoundDriver
	}
	return
}

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
