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

const createAccount = `-- name: CreateAccount :exec
INSERT INTO accounts (id, users, auth_info, wallet, level) VALUES ($1, $2, $3, $4, $5)
`

type CreateAccountParams struct {
	ID       uuid.UUID `json:"id"`
	Users    uuid.UUID `json:"users"`
	AuthInfo uuid.UUID `json:"auth_info"`
	Wallet   uuid.UUID `json:"wallet"`
	Level    uuid.UUID `json:"level"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount,
		arg.ID,
		arg.Users,
		arg.AuthInfo,
		arg.Wallet,
		arg.Level,
	)
	return err
}

const createAuthInfo = `-- name: CreateAuthInfo :exec
INSERT INTO auth_info (id, registered_number, pin) VALUES ($1, $2, $3)
`

type CreateAuthInfoParams struct {
	ID               uuid.UUID `json:"id"`
	RegisteredNumber int32     `json:"registered_number"`
	Pin              string    `json:"pin"`
}

func (q *Queries) CreateAuthInfo(ctx context.Context, arg CreateAuthInfoParams) error {
	_, err := q.db.ExecContext(ctx, createAuthInfo, arg.ID, arg.RegisteredNumber, arg.Pin)
	return err
}

const createLevel = `-- name: CreateLevel :exec
INSERT INTO levels (id, name) VALUES ($1, $2)
`

type CreateLevelParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) CreateLevel(ctx context.Context, arg CreateLevelParams) error {
	_, err := q.db.ExecContext(ctx, createLevel, arg.ID, arg.Name)
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

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, firstname, lastname, created_at, email, birth, phone) VALUES ($1, $2, $3, DEFAULT, $4, $5, $6)
`

type CreateUserParams struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Birth     time.Time `json:"birth"`
	Phone     string    `json:"phone"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Birth,
		arg.Phone,
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

const getLevelID = `-- name: GetLevelID :one
SELECT id FROM levels WHERE name = $1
`

func (q *Queries) GetLevelID(ctx context.Context, name string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getLevelID, name)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserWallet = `-- name: GetUserWallet :one
SELECT a.wallet FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1
`

func (q *Queries) GetUserWallet(ctx context.Context, registeredNumber int32) (uuid.NullUUID, error) {
	row := q.db.QueryRowContext(ctx, getUserWallet, registeredNumber)
	var wallet uuid.NullUUID
	err := row.Scan(&wallet)
	return wallet, err
}

const getUserWalletAndAuth = `-- name: GetUserWalletAndAuth :one
SELECT a.wallet, ai.pin FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1
`

type GetUserWalletAndAuthRow struct {
	Wallet uuid.NullUUID `json:"wallet"`
	Pin    string        `json:"pin"`
}

func (q *Queries) GetUserWalletAndAuth(ctx context.Context, registeredNumber int32) (GetUserWalletAndAuthRow, error) {
	row := q.db.QueryRowContext(ctx, getUserWalletAndAuth, registeredNumber)
	var i GetUserWalletAndAuthRow
	err := row.Scan(&i.Wallet, &i.Pin)
	return i, err
}
