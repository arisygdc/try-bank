package pgrepo

import (
	"context"
	"try-bank/database/postgres"
	"try-bank/request"
	"try-bank/util"

	"github.com/google/uuid"
)

func (d DB) CreateCompany(ctx context.Context, req request.PostCompany) (int32, error) {
	company := postgres.CreateCompanyParams{
		ID:    uuid.New(),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	wallet := postgres.CreateWalletParams{
		ID:      uuid.New(),
		Balance: req.TopUp,
	}

	authInfo := postgres.CreateAuthInfoParams{
		ID:               uuid.New(),
		RegisteredNumber: int32(req.Phone[9]+req.Phone[10]+req.Phone[11]) + int32(util.RandNum(5)),
		Pin:              req.Pin,
	}

	companyAccount := postgres.CreateCompanyAccountParams{
		ID:       uuid.New(),
		Company:  company.ID,
		AuthInfo: authInfo.ID,
		Wallet:   wallet.ID,
	}

	return authInfo.RegisteredNumber, d.transaction(ctx, func(q *postgres.Queries) error {
		if err := q.CreateCompany(ctx, company); err != nil {
			return err
		}

		if err := q.CreateWallet(ctx, wallet); err != nil {
			return err
		}

		if err := q.CreateAuthInfo(ctx, authInfo); err != nil {
			return err
		}

		if err := q.CreateCompanyAccount(ctx, companyAccount); err != nil {
			return err
		}
		return nil
	})
}

// Return virtual account identity, va key, error
func (d DB) ActivateVA(ctx context.Context, req request.VirtualAccount) (int32, string, error) {
	validateComp := postgres.ValidateCompanyParams{
		Name:             req.Name,
		Email:            req.Email,
		Phone:            req.Phone,
		RegisteredNumber: req.RegNum,
	}

	virtualAccount := postgres.CreateVirtualAccountParams{
		ID:                uuid.New(),
		VaKey:             util.RandString(32),
		FqdnDetailPayment: req.FQDNCheck,
		FqdnPay:           req.FQDNPay,
	}

	accountVA := postgres.UpdateVAstatusParams{
		VirtualAccount: uuid.NullUUID{
			UUID:  virtualAccount.ID,
			Valid: true,
		},
	}

	if err := d.transaction(ctx, func(q *postgres.Queries) error {
		vaID, err := q.ValidateCompany(ctx, validateComp)
		if err != nil {
			return err
		}

		accountVA.ID = vaID.UUID
		if err := q.CreateVirtualAccount(ctx, virtualAccount); err != nil {
			return err
		}

		if err := q.UpdateVAstatus(ctx, accountVA); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, "", err
	}

	vaIdentity, err := d.queries.GetVAIdentity(ctx, virtualAccount.ID)
	return vaIdentity, virtualAccount.VaKey, err
}
