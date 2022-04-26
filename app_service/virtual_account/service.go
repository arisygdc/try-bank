package virtualaccount

import (
	"context"
	"try-bank/database"
)

type ISVirtualAccount interface {
	Register(ctx context.Context, callback_url string) (RegistertrationVirtualAccountDetail, error)
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) Service {
	return Service{
		repos: repository,
	}
}
