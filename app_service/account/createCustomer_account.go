package account

import (
	"context"
	"fmt"
	"time"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type CreateCostumerParam struct {
	Firstname, Lastname, Email, Phone, Pin string
	Birth                                  time.Time
	TopUp                                  float64
	AccountType                            uuid.UUID
}

type RegisterCostumerDetail struct {
	Name, Email, Phone string
	Birth              time.Time
	TopUp              float64
	RegisteredNumber   int32
}

// create customer account is collection of inserting in table customer, auth_info, wallet, and account
// then provide inserted cusotmer information and error

// AccountType uuid can be filled by account.GetAccountType()
// use util.StrToTime() to validate birth
func (svc Service) CreateCustomerAccount(ctx context.Context, param CreateCostumerParam) (RegisterCostumerDetail, error) {
	var detailCostumer = RegisterCostumerDetail{}

	customer_param := postgresql.CreateCustomerParams{
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
		ID:            uuid.New(),
		CutomerID:     customer_param.ID,
		AuthInfoID:    authInfo_param.ID,
		WalletID:      wallet_param.ID,
		AccountTypeID: param.AccountType,
	}

	err := svc.repos.QueryTx(ctx, func(q *postgresql.Queries) error {
		err := q.CreateCustomer(ctx, customer_param)
		if err != nil {
			return err
		}

		detailCostumer.RegisteredNumber, err = q.CreateAuthInfo(ctx, authInfo_param)

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
		return detailCostumer, err
	}

	detailCostumer.Name = fmt.Sprintf("%s %s", customer_param.Firstname, customer_param.Lastname)
	detailCostumer.Email = param.Email
	detailCostumer.Birth = param.Birth
	detailCostumer.TopUp = param.TopUp
	detailCostumer.Phone = param.Phone

	return detailCostumer, err
}
