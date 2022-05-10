package account

import (
	"context"
	"errors"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

// Transfer balance from customer wallet to other customer wallet
func (svc Service) Transfer(ctx context.Context, fromWallet, toWallet uuid.UUID, balance float64) error {
	return svc.repos.QueryTx(ctx, func(q *postgresql.Queries) error {
		updated, err := q.SubtractBalance(ctx, postgresql.SubtractBalanceParams{
			ID:      toWallet,
			Balance: balance,
		})

		if updatedCheck(updated, err) != nil {
			return err
		}

		updated, err = q.AddBalance(ctx, postgresql.AddBalanceParams{
			ID:      fromWallet,
			Balance: balance,
		})

		if updatedCheck(updated, err) != nil {
			return err
		}

		return q.Transfer(ctx, postgresql.TransferParams{
			ID:         uuid.New(),
			FromWallet: fromWallet,
			ToWallet:   toWallet,
			Balance:    balance,
		})
	})
}

// check if only one row updated
func updatedCheck(updated int64, err error) error {
	if err != nil {
		return err
	}

	if updated < 1 {
		return errors.New("no rows updated")
	}

	if updated > 1 {
		return errors.New("too many rows updated")
	}

	return nil
}
