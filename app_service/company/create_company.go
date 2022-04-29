package company

import (
	"context"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type PublicInfo_company struct {
	Name, Email, Phone string
}

// PublicInfo_company contains string of Name, Email, Phone
type RegisterCompanyParam struct {
	PublicInfo_company
	Pin     string
	Deposit float64
}

type RegisteredCompanyDetail struct {
	PublicInfo_company
	Deposit        float64
	RegisterNumber int32
}

type CompaniesAccount struct {
	CompanyID        uuid.UUID
	AuthInfoID       uuid.UUID
	WalletID         uuid.UUID
	VirtualAccountID uuid.NullUUID
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
	return registered, err
}

func (svc Service) CompanyAccount(ctx context.Context, regNum_comp int32) (CompaniesAccount, error) {
	ca, err := svc.repos.Query().AuthGetCompaniesAccount(ctx, regNum_comp)
	return CompaniesAccount(ca), err
}
