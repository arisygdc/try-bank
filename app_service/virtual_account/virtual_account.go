package virtualaccount

import (
	"context"
	"strconv"

	"github.com/google/uuid"
)

type (
	virtualAccountIdentity int32
	virtualAccountNumber   int32
)

type CompaniesAccount struct {
	CompanyID        uuid.UUID
	AuthInfoID       uuid.UUID
	WalletID         uuid.UUID
	VirtualAccountID uuid.UUID
}

func (svc Service) VirtualAccountGetCompany(ctx context.Context, va_identity int32) (CompaniesAccount, error) {
	ca, err := svc.repos.Query().VAGetCompaniesAccount(ctx, va_identity)
	return CompaniesAccount{
		CompanyID:        ca.CompanyID,
		AuthInfoID:       ca.AuthInfoID,
		WalletID:         ca.WalletID,
		VirtualAccountID: ca.VirtualAccountID.UUID,
	}, err
}

func (svc Service) ValidateVirtualAccount(virtual_account string) (virtualAccountIdentity, virtualAccountNumber, error) {
	leng := len(virtual_account)
	a, b := virtual_account[:leng-3], virtual_account[3:]

	identity, err := strconv.Atoi(a)
	if err != nil {
		return 0, 0, err
	}

	number, err := strconv.Atoi(b)
	if err != nil {
		return 0, 0, err
	}

	return virtualAccountIdentity(identity), virtualAccountNumber(number), nil
}
