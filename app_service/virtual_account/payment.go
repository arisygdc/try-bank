package virtualaccount

import (
	"context"
	"errors"
	"time"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type PaidVA struct {
	ID              uuid.UUID
	IssuedPaymentID uuid.UUID
	PaidAt          time.Time
}

type PayVA struct {
	IssuedPayment uuid.UUID
	PayerWallet   uuid.UUID
	OwnerVAWallet uuid.UUID
	PaymentCharge float64
}

// TODO
// notify company using callback_url
func (svc Service) PaymentVirtualAccount(ctx context.Context, param PayVA) (PaidVA, error) {
	var paid PaidVA
	err := svc.repos.QueryTx(ctx, func(q *postgresql.Queries) error {
		payment, err := q.PaymentVA(ctx, postgresql.PaymentVAParams{
			ID:              uuid.New(),
			IssuedPaymentID: param.IssuedPayment,
		})
		if err != nil {
			return err
		}

		changes, err := q.SubtractBalance(ctx, postgresql.SubtractBalanceParams{
			ID:      param.PayerWallet,
			Balance: param.PaymentCharge,
		})

		if err != nil {
			return err
		}

		if changes < 1 {
			return errors.New("no rows changed")
		}

		changes, err = q.AddBalance(ctx, postgresql.AddBalanceParams{
			ID:      param.OwnerVAWallet,
			Balance: param.PaymentCharge,
		})

		if err != nil {
			return err
		}

		if changes < 1 {
			return errors.New("no rows changed")
		}

		paid = PaidVA(payment)
		return nil
	})

	return paid, err
}
