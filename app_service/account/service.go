package account

import (
	"context"
	"try-bank/database"

	"github.com/google/uuid"
)

const (
	LevelSilver   = "silver"
	LevelGold     = "gold"
	LevelPlatinum = "platinum"
)

type ISAccount interface {
	CreateCustomerAccount(ctx context.Context, param CreateCostumerParam) (RegisterCostumerDetail, error)
	GetAccountType(ctx context.Context, name_level string) (AccountType, error)
	CustomerAccount(ctx context.Context, regNum_comp int32) (CustomerAccount, error)
	Transfer(ctx context.Context, fromWallet, toWallet uuid.UUID, balance float64) error
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) ISAccount {
	return Service{
		repos: repository,
	}
}
