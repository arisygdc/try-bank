package virtualaccount

import "try-bank/database"

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) Service {
	return Service{
		repos: repository,
	}
}
