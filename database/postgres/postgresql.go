package postgres

import (
	"context"
	"database/sql"
	"strings"
	"time"
	"try-bank/request"
	"try-bank/util"

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
	authInfo := CreateAuthInfoParams{
		ID:               uuid.New(),
		RegisteredNumber: int32(t.Month()) + int32(req.Phone[9]+req.Phone[10]+req.Phone[11]),
		Pin:              req.Pin,
	}
	wallet := CreateWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	return d.transaction(ctx, func(query *Queries) error {
		if err := query.CreateUser(ctx, user); err != nil {
			return err
		}

		if err := query.CreateAuthInfo(ctx, authInfo); err != nil {
			return err
		}

		if err := query.CreateWallet(ctx, wallet); err != nil {
			return err
		}

		permID, err := query.GetLevelID(ctx, permission)
		if err != nil {
			return err
		}

		err = query.CreateAccount(ctx, CreateAccountParams{
			ID:       uuid.New(),
			Users:    user.ID,
			AuthInfo: authInfo.ID,
			Wallet:   wallet.ID,
			Level:    permID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (d DB) CreateCompany(ctx context.Context, req request.PostCompany) error {
	company := CreateCompanyParams{
		ID:    uuid.New(),
		Name:  req.Name,
		Email: req.Email,
	}

	wallet := CreateWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	authInfo := CreateAuthInfoParams{
		ID:               uuid.New(),
		RegisteredNumber: int32(req.Phone[9] + req.Phone[10] + req.Phone[11]),
		Pin:              req.Pin,
	}

	companyAccount := CreateCompanyAccountParams{
		ID:       uuid.New(),
		Company:  company.ID,
		AuthInfo: authInfo.ID,
		Wallet:   wallet.ID,
	}

	return d.transaction(ctx, func(q *Queries) error {
		if err := q.CreateCompany(ctx, company); err != nil {
			return err
		}

		if err := q.CreateWallet(ctx, wallet); err != nil {
			return err
		}

		if err := q.CreateAuthInfo(ctx, authInfo); err != nil {
			return err
		}

		if err := q.CreateCompanyAccount(ctx, companyAccount); err != nil {
			return err
		}
		return nil
	})
}

func (d DB) ActivateVA(ctx context.Context, req request.VirtualAccount) error {
	validateComp := ValidateCompanyParams{}
	virtualAccount := CreateVirtualAccountParams{
		ID:     uuid.New(),
		VaKey:  util.RandString(32),
		Domain: req.Domain,
	}

	accountVA := UpdateVAstatusParams{
		VirtualAccount: uuid.NullUUID{
			UUID:  virtualAccount.ID,
			Valid: true,
		},
	}

	return d.transaction(ctx, func(q *Queries) error {
		vaID, err := q.ValidateCompany(ctx, validateComp)
		if err != nil {
			return err
		}

		accountVA.ID = vaID.UUID
		if err := q.CreateVirtualAccount(ctx, virtualAccount); err != nil {
			return err
		}

		if err := q.UpdateVAstatus(ctx, accountVA); err != nil {
			return err
		}
		return nil
	})
}
