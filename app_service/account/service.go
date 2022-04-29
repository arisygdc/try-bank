package account

import (
	"context"
	"try-bank/database"
)

const (
	LevelSilver   = "silver"
	LevelGold     = "gold"
	LevelPlatinum = "platinum"
)

type ISAccount interface {
	CreateCustomerAccount(ctx context.Context, param CreateCostumerParam) (RegisterCostumerDetail, error)
	GetAccountType(ctx context.Context, name_level string) (AccountType, error)
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) ISAccount {
	return Service{
		repos: repository,
	}
}
