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
	ID uuid.UUID
	IssueVAPayment
	IssuedAt time.Time
}

func (svc Service) IssueVAPayment(ctx context.Context, param IssueVAPayment) error {
	return svc.repos.Query().IssuePaymentVA(ctx, postgresql.IssuePaymentVAParams{
		ID:                   uuid.New(),
		VirtualAccountID:     param.Virtual_account_id,
		VirtualAccountNumber: param.Virtual_account_number,
		PaymentCharge:        param.Payment_charge,
	})
}

func (svc Service) IssuedVAPaymentValidation(ctx context.Context, param IssueVAPayment) (IssuedPayment, error) {
	issuedPaymentParam := postgresql.CheckIssuedPaymentVAParams{
		VirtualAccountID:     param.Virtual_account_id,
		VirtualAccountNumber: param.Virtual_account_number,
	}

	issuedPayment, err := svc.repos.Query().CheckIssuedPaymentVA(ctx, issuedPaymentParam)
	if err != nil {
		return IssuedPayment{}, err
	}

	if issuedPayment.PaymentCharge != param.Payment_charge {
		return IssuedPayment{}, errors.New("the amount paid does not match")
	}

	return IssuedPayment{
		ID: issuedPayment.ID,
		IssueVAPayment: IssueVAPayment{
			Virtual_account_id:     issuedPayment.VirtualAccountID,
			Virtual_account_number: issuedPayment.VirtualAccountNumber,
			Payment_charge:         issuedPayment.PaymentCharge,
		},
		IssuedAt: issuedPayment.IssuedAt,
	}, nil
}
