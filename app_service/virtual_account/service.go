package virtualaccount

import (
	"context"
	"try-bank/database"

	"github.com/google/uuid"
)

type ISVirtualAccount interface {
	Register(ctx context.Context, company_id uuid.UUID, callback_url string) (RegistertrationVirtualAccountDetail, error)
	IssueVAPayment(ctx context.Context, param IssueVAPayment) error
	IssuedVAPaymentValidation(ctx context.Context, param IssueVAPayment) (IssuedPayment, error)
	ValidateVirtualAccount(virtual_account string) (virtualAccountIdentity, virtualAccountNumber, error)
	PaymentVirtualAccount(ctx context.Context, param PayVA) (PaidVA, error)
	VirtualAccountGetCompany(ctx context.Context, va_identity int32) (CompaniesAccount, error)
}

type Service struct {
	repos database.IRepository
}

func New(repository database.IRepository) ISVirtualAccount {
	return Service{
		repos: repository,
	}
}
