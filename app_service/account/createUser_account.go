package account

import (
	"context"
	"fmt"
	"time"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type CreateUserParam struct {
	Firstname, Lastname, Email, Phone, Pin string
	Birth                                  time.Time
	TopUp                                  float64
	Level                                  string
}

type RegisterUserDetail struct {
	Name, Email, Phone string
	Birth              time.Time
	TopUp              float64
	RegisteredNumber   int32
}

func (svc Service) CreateUserAccount(ctx context.Context, param CreateUserParam) (RegisterUserDetail, error) {
	var detailUser = RegisterUserDetail{}

	lvl, err := svc.GetLevel(ctx, param.Level)
	if err != nil {
		return detailUser, err
	}

	// Still handle create client user only
	if lvl.Name != LevelClientUser {
		return detailUser, err
	}

	user_param := postgresql.CreateUserParams{
		ID:        uuid.New(),
		Firstname: param.Firstname,
		Lastname:  param.Lastname,
		Email:     param.Email,
		Birth:     param.Birth,
		Phone:     param.Phone,
	}

	authInfo_param := postgresql.CreateAuthInfoParams{
		ID:  uuid.New(),
		Pin: param.Pin,
	}

	wallet_param := postgresql.CreateWalletParams{
		ID:      uuid.New(),
		Balance: param.TopUp,
	}

	account := postgresql.CreateAccountParams{
		ID:       uuid.New(),
		Users:    user_param.ID,
		AuthInfo: authInfo_param.ID,
		Wallet:   wallet_param.ID,
		Level:    lvl.ID,
	}

	err = svc.repos.QueryTx(ctx, func(q *postgresql.Queries) error {
		err := q.CreateUser(ctx, user_param)
		if err != nil {
			return err
		}

		detailUser.RegisteredNumber, err = q.CreateAuthInfo(ctx, authInfo_param)

		if err != nil {
			return err
		}

		err = q.CreateWallet(ctx, wallet_param)

		if err != nil {
			return err
		}

		err = q.CreateAccount(ctx, account)
		return err
	})

	if err != nil {
		return detailUser, err
	}

	detailUser.Name = fmt.Sprintf("%s %s", user_param.Firstname, user_param.Lastname)
	detailUser.Email = param.Email
	detailUser.Birth = param.Birth
	detailUser.TopUp = param.TopUp
	detailUser.Phone = param.Phone

	return detailUser, err
}
