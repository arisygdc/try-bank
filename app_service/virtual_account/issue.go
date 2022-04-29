package virtualaccount

import (
	"context"
	"errors"
	"time"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type IssueVAPayment struct {
	Virtual_account_id     uuid.UUID
	Virtual_account_number int32
	Payment_charge         float64
}

type IssuedPayment struct {
	ID                   uuid.UUID
	VirtualAccountID     uuid.UUID
	VirtualAccountNumber int32
	PaymentCharge        float64
	IssuedAt             time.Time
}

func (svc Service) IssueVAPayment(ctx context.Context, param IssueVAPayment) error {
	return svc.repos.Query().IssuePaymentVA(ctx, postgresql.IssuePaymentVAParams{
		ID:                   uuid.New(),
		VirtualAccountID:     param.Virtual_account_id,
		VirtualAccountNumber: param.Virtual_account_number,
		PaymentCharge:        param.Payment_charge,
	})
}

func (svc Service) CheckIssuedVAPayment(ctx context.Context, virtualAccount_id uuid.UUID, virtualAccount_number int32, amount float64) (IssuedPayment, error) {
	issuedPaymentParam := postgresql.CheckIssuedPaymentVAParams{
		VirtualAccountID:     virtualAccount_id,
		VirtualAccountNumber: virtualAccount_number,
	}

	issuedPayment, err := svc.repos.Query().CheckIssuedPaymentVA(ctx, issuedPaymentParam)
	if err != nil {
		return IssuedPayment{}, err
	}

	if issuedPayment.PaymentCharge != amount {
		return IssuedPayment{}, errors.New("the amount paid does not match")
	}

	return IssuedPayment(issuedPayment), nil
}
