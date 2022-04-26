// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: account.sql

package postgresql

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const accountType = `-- name: AccountType :one
SELECT id, max_transfer, name FROM account_type WHERE name = $1
`

func (q *Queries) AccountType(ctx context.Context, name string) (AccountType, error) {
	row := q.db.QueryRowContext(ctx, accountType, name)
	var i AccountType
	err := row.Scan(&i.ID, &i.MaxTransfer, &i.Name)
	return i, err
}

const createAccount = `-- name: CreateAccount :exec
INSERT INTO accounts (id, cutomer_id, auth_info_id, wallet_id, account_type_id) VALUES ($1, $2, $3, $4, $5)
`

type CreateAccountParams struct {
	ID            uuid.UUID `json:"id"`
	CutomerID     uuid.UUID `json:"cutomer_id"`
	AuthInfoID    uuid.UUID `json:"auth_info_id"`
	WalletID      uuid.UUID `json:"wallet_id"`
	AccountTypeID uuid.UUID `json:"account_type_id"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount,
		arg.ID,
		arg.CutomerID,
		arg.AuthInfoID,
		arg.WalletID,
		arg.AccountTypeID,
	)
	return err
}

const createAuthInfo = `-- name: CreateAuthInfo :one
INSERT INTO auth_info (id, registered_number, pin) VALUES ($1, DEFAULT, $2) RETURNING registered_number
`

type CreateAuthInfoParams struct {
	ID  uuid.UUID `json:"id"`
	Pin string    `json:"pin"`
}

func (q *Queries) CreateAuthInfo(ctx context.Context, arg CreateAuthInfoParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createAuthInfo, arg.ID, arg.Pin)
	var registered_number int32
	err := row.Scan(&registered_number)
	return registered_number, err
}

const createCustomer = `-- name: CreateCustomer :exec
INSERT INTO cutomers (id, firstname, lastname, created_at, email, birth, phone) VALUES ($1, $2, $3, DEFAULT, $4, $5, $6)
`

type CreateCustomerParams struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Birth     time.Time `json:"birth"`
	Phone     string    `json:"phone"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) error {
	_, err := q.db.ExecContext(ctx, createCustomer,
		arg.ID,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Birth,
		arg.Phone,
	)
	return err
}

const createTransfer = `-- name: CreateTransfer :exec
INSERT INTO transfers (id, from_wallet, to_wallet, balance, transfer_at) VALUES ($1, $2, $3, $4, DEFAULT)
`

type CreateTransferParams struct {
	ID         uuid.UUID `json:"id"`
	FromWallet uuid.UUID `json:"from_wallet"`
	ToWallet   uuid.UUID `json:"to_wallet"`
	Balance    float64   `json:"balance"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) error {
	_, err := q.db.ExecContext(ctx, createTransfer,
		arg.ID,
		arg.FromWallet,
		arg.ToWallet,
		arg.Balance,
	)
	return err
}

const createWallet = `-- name: CreateWallet :exec
INSERT INTO wallets (id, balance, last_update) VALUES ($1, $2, DEFAULT)
`

type CreateWalletParams struct {
	ID      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}

func (q *Queries) CreateWallet(ctx context.Context, arg CreateWalletParams) error {
	_, err := q.db.ExecContext(ctx, createWallet, arg.ID, arg.Balance)
	return err
}

const getUserWallet = `-- name: GetUserWallet :one
SELECT a.wallet_id FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1
`

func (q *Queries) GetUserWallet(ctx context.Context, registeredNumber int32) (uuid.NullUUID, error) {
	row := q.db.QueryRowContext(ctx, getUserWallet, registeredNumber)
	var wallet_id uuid.NullUUID
	err := row.Scan(&wallet_id)
	return wallet_id, err
}

const getUserWalletAndAuth = `-- name: GetUserWalletAndAuth :one
SELECT a.wallet_id, ai.pin FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1
`

type GetUserWalletAndAuthRow struct {
	WalletID uuid.NullUUID `json:"wallet_id"`
	Pin      string        `json:"pin"`
}

func (q *Queries) GetUserWalletAndAuth(ctx context.Context, registeredNumber int32) (GetUserWalletAndAuthRow, error) {
	row := q.db.QueryRowContext(ctx, getUserWalletAndAuth, registeredNumber)
	var i GetUserWalletAndAuthRow
	err := row.Scan(&i.WalletID, &i.Pin)
	return i, err
}
