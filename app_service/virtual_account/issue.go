package virtualaccount

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"try-bank/database/postgresql"

	"github.com/google/uuid"
)

var ErrRepeatedIssue = fmt.Errorf("cannot repeat issue")

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

// check available issue
func (svc Service) IssueVAPayment(ctx context.Context, param IssueVAPayment) error {
	_, err := svc.repos.Query().CheckActiveIssueVAP(ctx, postgresql.CheckActiveIssueVAPParams{
		VirtualAccountID:     param.Virtual_account_id,
		VirtualAccountNumber: param.Virtual_account_number,
	})

	if err != sql.ErrNoRows {
		return ErrRepeatedIssue
	}

	return svc.repos.Query().IssuePaymentVA(ctx, postgresql.IssuePaymentVAParams{
		ID:                   uuid.New(),
		VirtualAccountID:     param.Virtual_account_id,
		VirtualAccountNumber: param.Virtual_account_number,
		PaymentCharge:        param.Payment_charge,
	})
}

func (svc Service) IssuedVAPaymentValidation(ctx context.Context, param IssueVAPayment) (IssuedPayment, error) {
	issuedPaymentParam := postgresql.CheckActiveIssueVAPParams{
		VirtualAccountID:     param.Virtual_account_id,
		VirtualAccountNumber: param.Virtual_account_number,
	}

	issuedPayment, err := svc.repos.Query().CheckActiveIssueVAP(ctx, issuedPaymentParam)
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
