package pgrepo

import (
	"context"
	"strings"
	"time"
	"try-bank/database/postgres"
	"try-bank/request"

	"github.com/google/uuid"
)

func (d DB) CreateLevel(ctx context.Context, req request.PermissionReq) error {
	return d.queries.CreateLevel(ctx, postgres.CreateLevelParams{
		ID:   uuid.New(),
		Name: req.Name,
	})
}

func (d DB) CreateUser(ctx context.Context, req request.PostUser, permission string) error {
	t, err := time.Parse("2006-1-2", strings.Trim(req.Birth, " "))
	if err != nil {
		return err
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
		RegisteredNumber: int32(t.Month()) + int32(req.Phone[9]+req.Phone[10]+req.Phone[11]),
		Pin:              req.Pin,
	}

	wallet := postgres.CreateWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	return d.transaction(ctx, func(query *postgres.Queries) error {
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
