package company

import (
	"context"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type RegisterCompanyParam struct {
	Name, Email, Phone, Pin string
	Deposit                 float64
}

type RegisteredCompanyDetail struct {
	Name, Email, Phone                                string
	Deposit                                           float64
	RegisterNumber                                    int32
	IdCompany, IdAuthInfo, IdWallet, IdCompanyAccount string
}

func (svc Service) CreateCompanyAccount(ctx context.Context, param RegisterCompanyParam) (RegisteredCompanyDetail, error) {
	var registered = RegisteredCompanyDetail{}

	company := postgresql.CreateCompanyParams{
		ID:    uuid.New(),
		Name:  param.Name,
		Email: param.Email,
		Phone: param.Phone,
	}

	auth_info := postgresql.CreateAuthInfoParams{
		ID:  uuid.New(),
		Pin: param.Pin,
	}

	wallet := postgresql.CreateWalletParams{
		ID:      uuid.New(),
		Balance: param.Deposit,
	}

	companyAccount := postgresql.CreateCompanyAccountParams{
		ID:         uuid.New(),
		CompanyID:  company.ID,
		AuthInfoID: auth_info.ID,
		WalletID:   wallet.ID,
	}

	err := svc.repos.QueryTx(ctx, func(q *postgresql.Queries) error {
		err := q.CreateCompany(ctx, company)
		if err != nil {
			return err
		}

		err = q.CreateWallet(ctx, wallet)
		if err != nil {
			return err
		}

		registered.RegisterNumber, err = q.CreateAuthInfo(ctx, auth_info)
		if err != nil {
			return err
		}

		err = q.CreateCompanyAccount(ctx, companyAccount)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return registered, err
	}

	registered.Name = company.Name
	registered.Email = company.Email
	registered.Phone = company.Phone
	registered.Deposit = wallet.Balance
	registered.IdCompany = company.ID.String()
	registered.IdAuthInfo = auth_info.ID.String()
	registered.IdWallet = wallet.ID.String()
	registered.IdCompanyAccount = companyAccount.ID.String()
	return registered, err
}
