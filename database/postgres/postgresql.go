package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
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

	database = &DB{
		Conn:    sqlconn,
		Queries: New(sqlconn),
	}
	return
}

func (d DB) transaction(ctx context.Context, queryFunc func(*Queries) error) error {
	tx, err := d.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queriesTx := d.Queries.WithTx(tx)
	err = queryFunc(queriesTx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

func (d DB) CreateUser(ctx context.Context, userParam CreateUserParams, authInfoParam CreateAuthInfoParams, walletParam CreateCoustomerWalletParams, permission string) error {
	return d.transaction(ctx, func(query *Queries) error {
		if err := query.CreateUser(ctx, userParam); err != nil {
			return err
		}

		if err := query.CreateAuthInfo(ctx, authInfoParam); err != nil {
			return err
		}

		if err := query.CreateCoustomerWallet(ctx, walletParam); err != nil {
			return err
		}

		permID, err := query.GetPermissionID(ctx, permission)
		if err != nil {
			return err
		}

		if err := query.CreateAccount(ctx, CreateAccountParams{
			ID:         uuid.New(),
			Users:      userParam.ID,
			AuthInfo:   authInfoParam.ID,
			Wallet:     walletParam.ID,
			Permission: permID,
		}); err != nil {
			return err
		}

		return nil
	})
}
