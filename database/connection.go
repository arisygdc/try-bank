package database

import (
	"context"
	"errors"
	"try-bank/config"
	"try-bank/database/postgres/pgrepo"
	"try-bank/request"
)

var (
	ErrNotFoundDriver = errors.New("database driver not found")
)

type IRepo interface {
	CreateLevel(ctx context.Context, req request.PermissionReq) error
	CreateUser(ctx context.Context, req request.PostUser, permission string) error
	CreateCompany(ctx context.Context, req request.PostCompany) error
	ActivateVA(ctx context.Context, req request.VirtualAccount) error
	PaymentVA(ctx context.Context, req request.PaymentVA) error
	Transfer(ctx context.Context, req request.Transfer) error
}

func NewSQL(env config.Environment) (dbsource IRepo, err error) {
	if env.DBDriver == "postgres" {
		dbsource, err = pgrepo.NewPostgres(env.DBDriver, env.DBSource)
	} else {
		err = ErrNotFoundDriver
	}
	return
}
