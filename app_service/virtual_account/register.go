package virtualaccount

import (
	"context"
	"crypto/rand"
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

	virtualAccountParam := virtualAccount_param(callback_url)

	setCompanyVaParam := setCompanyVa_param(company_id, virtualAccountParam.ID)

	err := svc.repos.QueryTx(ctx, func(q *postgresql.Queries) error {
		err := q.SetCompanyVA(ctx, setCompanyVaParam)
		if err != nil {
			return err
		}
		err = q.CreateVirtualAccount(ctx, virtualAccountParam)
		return err
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

func virtualAccount_param(callback_url string) postgresql.CreateVirtualAccountParams {
	id := uuid.New()

	random, _ := rand.Int(rand.Reader, big.NewInt(999))

	identity := int32(random.Int64())

	authKey := uuid.NewSHA1(id, random.FillBytes(random.Bytes()))

	return postgresql.CreateVirtualAccountParams{
		ID:               id,
		Identity:         identity,
		AuthorizationKey: authKey.String(),
		CallbackUrl:      callback_url,
	}
}

func setCompanyVa_param(company_id uuid.UUID, virtualAccount_id uuid.UUID) postgresql.SetCompanyVAParams {
	return postgresql.SetCompanyVAParams{
		ID: company_id,
		VirtualAccountID: uuid.NullUUID{
			UUID:  virtualAccount_id,
			Valid: true,
		},
	}
}
