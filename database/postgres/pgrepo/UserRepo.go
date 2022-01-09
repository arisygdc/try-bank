package pgrepo

import (
	"context"
	"strings"
	"time"
	"try-bank/database/postgres"
	"try-bank/request"
	"try-bank/util"

	"github.com/google/uuid"
)

func (d DB) CreateLevel(ctx context.Context, req request.PermissionReq) error {
	return d.queries.CreateLevel(ctx, postgres.CreateLevelParams{
		ID:   uuid.New(),
		Name: req.Name,
	})
}

func (d DB) CreateUser(ctx context.Context, req request.PostUser, permission string) (int32, error) {
	t, err := time.Parse("2006-1-2", strings.Trim(req.Birth, " "))
	if err != nil {
		return 0, err
	}

	user := postgres.CreateUserParams{
		ID:        uuid.New(),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Birth:     t.UTC(),
		Phone:     req.Phone,
	}

	authInfo := postgres.CreateAuthInfoParams{
		ID:               uuid.New(),
		RegisteredNumber: int32(util.RandNum(5)),
		Pin:              req.Pin,
	}

	wallet := postgres.CreateWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	return authInfo.RegisteredNumber, d.transaction(ctx, func(query *postgres.Queries) error {
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

		err = query.CreateAccount(ctx, postgres.CreateAccountParams{
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

func (r DB) Login(ctx context.Context, regNum int32, pin string) (postgres.AuthInfo, error) {
	return r.queries.Login(ctx, postgres.LoginParams{
		RegisteredNumber: regNum,
		Pin:              pin,
	})
}

func (r DB) CekSaldo(ctx context.Context, regNum int32) (float64, error) {
	wallet, err := r.queries.GetUserWallet(ctx, regNum)
	if err != nil {
		return 0, err
	}
	return r.queries.GetBalance(ctx, wallet.UUID)
}

func (r DB) GetAccount(ctx context.Context, regNum int32) (uuid.UUID, error) {
	return r.queries.GetAccount(ctx, regNum)
}

func (r DB) GetUser(ctx context.Context, account uuid.UUID) (postgres.User, error) {
	return r.queries.GetUserInfo(ctx, account)
}
