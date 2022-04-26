// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: company.sql

package postgresql

import (
	"context"

	"github.com/google/uuid"
)

const createCompany = `-- name: CreateCompany :exec
INSERT INTO companies (id, name, email, phone, created_at) VALUES ($1, $2, $3, $4, DEFAULT)
`

type CreateCompanyParams struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) error {
	_, err := q.db.ExecContext(ctx, createCompany,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Phone,
	)
	return err
}

const createCompanyAccount = `-- name: CreateCompanyAccount :exec
INSERT INTO companies_account (id, company_id, auth_info_id, wallet_id, virtual_account_id) VALUES ($1, $2, $3, $4, $5)
`

type CreateCompanyAccountParams struct {
	ID               uuid.UUID     `json:"id"`
	CompanyID        uuid.UUID     `json:"company_id"`
	AuthInfoID       uuid.UUID     `json:"auth_info_id"`
	WalletID         uuid.UUID     `json:"wallet_id"`
	VirtualAccountID uuid.NullUUID `json:"virtual_account_id"`
}

func (q *Queries) CreateCompanyAccount(ctx context.Context, arg CreateCompanyAccountParams) error {
	_, err := q.db.ExecContext(ctx, createCompanyAccount,
		arg.ID,
		arg.CompanyID,
		arg.AuthInfoID,
		arg.WalletID,
		arg.VirtualAccountID,
	)
	return err
}

const createVAPayment = `-- name: CreateVAPayment :exec
INSERT INTO va_payment (id, issued_payment_id, paid_at) VALUES ($1, $2, DEFAULT)
`

type CreateVAPaymentParams struct {
	ID              uuid.UUID `json:"id"`
	IssuedPaymentID uuid.UUID `json:"issued_payment_id"`
}

func (q *Queries) CreateVAPayment(ctx context.Context, arg CreateVAPaymentParams) error {
	_, err := q.db.ExecContext(ctx, createVAPayment, arg.ID, arg.IssuedPaymentID)
	return err
}

const createVirtualAccount = `-- name: CreateVirtualAccount :exec
INSERT INTO virtual_account (id, authorization_key, identity, callback_url, created_at) VALUES ($1, $2, $3, $4, DEFAULT)
`

type CreateVirtualAccountParams struct {
	ID               uuid.UUID `json:"id"`
	AuthorizationKey string    `json:"authorization_key"`
	Identity         int32     `json:"identity"`
	CallbackUrl      string    `json:"callback_url"`
}

func (q *Queries) CreateVirtualAccount(ctx context.Context, arg CreateVirtualAccountParams) error {
	_, err := q.db.ExecContext(ctx, createVirtualAccount,
		arg.ID,
		arg.AuthorizationKey,
		arg.Identity,
		arg.CallbackUrl,
	)
	return err
}

const setCompanyVA = `-- name: SetCompanyVA :exec
UPDATE companies_account SET virtual_account_id = $1 WHERE id = $2
`

type SetCompanyVAParams struct {
	VirtualAccountID uuid.NullUUID `json:"virtual_account_id"`
	ID               uuid.UUID     `json:"id"`
}

func (q *Queries) SetCompanyVA(ctx context.Context, arg SetCompanyVAParams) error {
	_, err := q.db.ExecContext(ctx, setCompanyVA, arg.VirtualAccountID, arg.ID)
	return err
}

const validateCompany = `-- name: ValidateCompany :one
SELECT ca.id FROM companies_account ca 
RIGHT JOIN companies c ON ca.company = c.id 
RIGHT JOIN auth_info ai ON ca.auth_info = ai.id
WHERE c.name = $1 AND c.email = $2 AND c.phone = $3
AND ai.registered_number = $4
`

type ValidateCompanyParams struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	RegisteredNumber int32  `json:"registered_number"`
}

func (q *Queries) ValidateCompany(ctx context.Context, arg ValidateCompanyParams) (uuid.NullUUID, error) {
	row := q.db.QueryRowContext(ctx, validateCompany,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.RegisteredNumber,
	)
	var id uuid.NullUUID
	err := row.Scan(&id)
	return id, err
}
