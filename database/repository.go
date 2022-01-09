package database

import (
	"context"
	"errors"
	"try-bank/config"
	"try-bank/database/postgres"
	"try-bank/database/postgres/pgrepo"
	"try-bank/request"

	"github.com/google/uuid"
)

var (
	ErrNotFoundDriver = errors.New("database driver not found")
)

type IRepo interface {
	CreateLevel(ctx context.Context, req request.PermissionReq) error
	CreateUser(ctx context.Context, req request.PostUser, permission string) (int32, error)
	CreateCompany(ctx context.Context, req request.PostCompany) (int32, error)
	ActivateVA(ctx context.Context, req request.VirtualAccount) (int32, string, error)
	PaymentVA(ctx context.Context, req request.PaymentVA) error
	Transfer(ctx context.Context, req request.Transfer) error
	Login(ctx context.Context, regNum int32, pin string) (postgres.AuthInfo, error)
	CekSaldo(ctx context.Context, regNum int32) (float64, error)
	GetAccount(ctx context.Context, regNum int32) (uuid.UUID, error)
	GetUser(ctx context.Context, account uuid.UUID) (postgres.User, error)
}

func NewRepository(env config.Environment) (dbsource IRepo, err error) {
	if env.DBDriver == "postgres" {
		dbsource, err = pgrepo.NewPostgres(env.DBDriver, env.DBSource)
	} else {
		err = ErrNotFoundDriver
	}
	return
}
