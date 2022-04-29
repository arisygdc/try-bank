package virtualaccount

import (
	"context"
	"try-bank/database"

	"github.com/google/uuid"
)

type ISVirtualAccount interface {
	Register(ctx context.Context, company_id uuid.UUID, callback_url string) (RegistertrationVirtualAccountDetail, error)
	IssueVAPayment(ctx context.Context, param IssueVAPayment) error
	PaymentVirtualAccount(ctx context.Context, IssuedPayment_id uuid.UUID) (PaidVA, error)
	VirtualAccountID(ctx context.Context, va_identity int32) (uuid.UUID, error)
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) ISVirtualAccount {
	return Service{
		repos: repository,
	}
}
