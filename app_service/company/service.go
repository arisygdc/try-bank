package company

import (
	"context"
	"try-bank/database"
)

type ISCompany interface {
	CreateCompanyAccount(ctx context.Context, param RegisterCompanyParam) (RegisteredCompanyDetail, error)
	CompanyAccount(ctx context.Context, regNum_comp int32) (CompaniesAccount, error)
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) ISCompany {
	return Service{
		repos: repository,
	}
}
