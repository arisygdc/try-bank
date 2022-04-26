package account

import (
	"context"

	"github.com/google/uuid"
)

type AccountType struct {
	ID          uuid.UUID
	MaxTransfer float64
	Name        string
}

func (svc Service) GetAccountType(ctx context.Context, name_level string) (AccountType, error) {
	getAccountType, err := svc.repos.Query().AccountType(ctx, name_level)
	return AccountType(getAccountType), err
}
