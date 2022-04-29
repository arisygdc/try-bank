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
	CreateUserAccount(ctx context.Context, param CreateCostumerParam) (RegisterCostumerDetail, error)
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) Service {
	return Service{
		repos: repository,
	}
}
