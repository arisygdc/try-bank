package virtualaccount

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type PaidVA struct {
	ID              uuid.UUID
	IssuedPaymentID uuid.UUID
	PaidAt          time.Time
}

func (svc Service) PaymentVirtualAccount(ctx context.Context, IssuedPayment_id uuid.UUID) (PaidVA, error) {
	paid, err := svc.repos.Query().PaymentVA(ctx, IssuedPayment_id)
	return PaidVA(paid), err
}
