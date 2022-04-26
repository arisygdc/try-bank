package account

import "try-bank/database"

const (
	LevelAdmin   = "admin"
	LevelUser    = "user"
	LevelCompany = "company"
)

type ISAccount interface {
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) Service {
	return Service{
		repos: repository,
	}
}
