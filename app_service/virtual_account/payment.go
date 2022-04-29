package virtualaccount

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type PaidVA struct {
	ID              uuid.UUID
	IssuedPaymentID uuid.UUID
	PaidAt          time.Time
}

// TODO
// use transaction to subtract payer wallet and increase company wallet
// notify company using callback_url
func (svc Service) PaymentVirtualAccount(ctx context.Context, IssuedPayment_id uuid.UUID) (PaidVA, error) {
	paid, _ := svc.repos.Query().PaymentVA(ctx, IssuedPayment_id)
	return PaidVA(paid), errors.New("<- todo")
}
