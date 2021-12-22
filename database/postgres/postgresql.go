package postgres

import (
	"context"
	"database/sql"
	"strings"
	"time"
	"try-bank/request"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type DB struct {
	conn    *sql.DB
	Queries *Queries
}

func NewPostgres(dbdriver, dbsource string) (database *DB, err error) {
	sqlconn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		return
	}

	database = &DB{
		conn:    sqlconn,
		Queries: New(sqlconn),
	}
	return
}

func (d DB) transaction(ctx context.Context, queryFunc func(*Queries) error) error {
	tx, err := d.conn.BeginTx(ctx, nil)
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

func (d DB) CreateUser(ctx context.Context, req request.PostUser, permission string) error {
	return d.transaction(ctx, func(query *Queries) error {
		t, err := time.Parse("2006-1-2", strings.Trim(req.Birth, " "))
		if err != nil {
			return err
		}

		user := CreateUserParams{
			ID:        uuid.New(),
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			Email:     req.Email,
			Birth:     t.UTC(),
			Phone:     req.Phone,
		}
		if err := query.CreateUser(ctx, user); err != nil {
			return err
		}

		authInfo := CreateAuthInfoParams{
			ID:               uuid.New(),
			RegisteredNumber: int32(t.Month()) + int32(req.Phone[9]+req.Phone[10]+req.Phone[11]),
		}
		if err := query.CreateAuthInfo(ctx, authInfo); err != nil {
			return err
		}

		wallet := CreateCoustomerWalletParams{
			ID:      uuid.New(),
			Balance: req.TopUp,
		}
		if err := query.CreateCoustomerWallet(ctx, wallet); err != nil {
			return err
		}

		permID, err := query.GetPermissionID(ctx, permission)
		if err != nil {
			return err
		}

		err = query.CreateAccount(ctx, CreateAccountParams{
			ID:         uuid.New(),
			Users:      user.ID,
			AuthInfo:   authInfo.ID,
			Wallet:     wallet.ID,
			Permission: permID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
