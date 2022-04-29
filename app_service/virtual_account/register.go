package virtualaccount

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

type RegistertrationVirtualAccountDetail struct {
	Identity         int32
	AuthorizationKey string
	Callback_url     string
}

func (svc Service) Register(ctx context.Context, company_id uuid.UUID, callback_url string) (RegistertrationVirtualAccountDetail, error) {
	var outputDetail RegistertrationVirtualAccountDetail
	id := uuid.New()

	random, _ := rand.Int(rand.Reader, big.NewInt(999))

	identity := int32(random.Int64())

	authKey := uuid.NewSHA1(id, random.FillBytes(random.Bytes()))

	virtualAccountParam := postgresql.CreateVirtualAccountParams{
		ID:               id,
		Identity:         identity,
		AuthorizationKey: authKey.String(),
		CallbackUrl:      callback_url,
	}

	setCompanyVaParam := postgresql.ActivateVirtualAccountParams{
		CompanyID: company_id,
		VirtualAccountID: uuid.NullUUID{
			UUID:  virtualAccountParam.ID,
			Valid: true,
		},
	}

	err := svc.repos.QueryTx(ctx, func(q *postgresql.Queries) error {
		err := q.CreateVirtualAccount(ctx, virtualAccountParam)
		if err != nil {
			return err
		}

		rowUpdated, err := q.ActivateVirtualAccount(ctx, setCompanyVaParam)
		if err != nil {
			return err
		}

		if rowUpdated < 1 {
			return errors.New("no rows updated")
		}

		return nil
	})

	if err != nil {
		return outputDetail, err
	}

	outputDetail = RegistertrationVirtualAccountDetail{
		Identity:         virtualAccountParam.Identity,
		AuthorizationKey: virtualAccountParam.AuthorizationKey,
		Callback_url:     virtualAccountParam.CallbackUrl,
	}

	return outputDetail, err
}
