package company

import "try-bank/database"

type ISCompany interface {
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) Service {
	return Service{
		repos: repository,
	}
}
