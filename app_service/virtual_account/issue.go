package virtualaccount

import (
	"context"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type IssueVAPayment struct {
	Virtual_account_id     uuid.UUID
	Virtual_account_number int32
	Payment_charge         float64
}

func (svc Service) IssueVAPayment(ctx context.Context, param IssueVAPayment) error {
	return svc.repos.Query().IssuePaymentVA(ctx, postgresql.IssuePaymentVAParams{
		ID:                   uuid.New(),
		VirtualAccountID:     param.Virtual_account_id,
		VirtualAccountNumber: param.Virtual_account_number,
		PaymentCharge:        param.Payment_charge,
	})
}
