// Code generated by sqlc. DO NOT EDIT.
// source: company.sql

package postgres

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
INSERT INTO companies_account (id, company, auth_info, wallet, virtual_account) VALUES ($1, $2, $3, $4, $5)
`

type CreateCompanyAccountParams struct {
	ID             uuid.UUID     `json:"id"`
	Company        uuid.UUID     `json:"company"`
	AuthInfo       uuid.UUID     `json:"auth_info"`
	Wallet         uuid.UUID     `json:"wallet"`
	VirtualAccount uuid.NullUUID `json:"virtual_account"`
}

func (q *Queries) CreateCompanyAccount(ctx context.Context, arg CreateCompanyAccountParams) error {
	_, err := q.db.ExecContext(ctx, createCompanyAccount,
		arg.ID,
		arg.Company,
		arg.AuthInfo,
		arg.Wallet,
		arg.VirtualAccount,
	)
	return err
}

const createVAPayment = `-- name: CreateVAPayment :exec
INSERT INTO va_payment (id, virtual_account, va_number, request_payment, paid_at) VALUES (NULL, $1, $2, $3, DEFAULT)
`

type CreateVAPaymentParams struct {
	VirtualAccount uuid.UUID `json:"virtual_account"`
	VaNumber       string    `json:"va_number"`
	RequestPayment float64   `json:"request_payment"`
}

func (q *Queries) CreateVAPayment(ctx context.Context, arg CreateVAPaymentParams) error {
	_, err := q.db.ExecContext(ctx, createVAPayment, arg.VirtualAccount, arg.VaNumber, arg.RequestPayment)
	return err
}

const createVirtualAccount = `-- name: CreateVirtualAccount :exec
INSERT INTO virtual_account (id, va_key, fqdn_detail_payment, fqdn_pay, created_at) VALUES ($1, $2, $3, $4, DEFAULT)
`

type CreateVirtualAccountParams struct {
	ID                uuid.UUID `json:"id"`
	VaKey             string    `json:"va_key"`
	FqdnDetailPayment string    `json:"fqdn_detail_payment"`
	FqdnPay           string    `json:"fqdn_pay"`
}

func (q *Queries) CreateVirtualAccount(ctx context.Context, arg CreateVirtualAccountParams) error {
	_, err := q.db.ExecContext(ctx, createVirtualAccount,
		arg.ID,
		arg.VaKey,
		arg.FqdnDetailPayment,
		arg.FqdnPay,
	)
	return err
}

const updateVAstatus = `-- name: UpdateVAstatus :exec
UPDATE companies_account SET virtual_account = $1 WHERE id = $2
`

type UpdateVAstatusParams struct {
	VirtualAccount uuid.NullUUID `json:"virtual_account"`
	ID             uuid.UUID     `json:"id"`
}

func (q *Queries) UpdateVAstatus(ctx context.Context, arg UpdateVAstatusParams) error {
	_, err := q.db.ExecContext(ctx, updateVAstatus, arg.VirtualAccount, arg.ID)
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
