-- name: CreateLevel :exec
INSERT INTO levels (id, name) VALUES ($1, $2);

-- name: GetLevelID :one
SELECT id FROM levels WHERE name = $1;

-- name: CreateUser :exec
INSERT INTO users (id, firstname, lastname, created_at, email, birth, phone) VALUES ($1, $2, $3, DEFAULT, $4, $5, $6);

-- name: CreateAuthInfo :exec
INSERT INTO auth_info (id, registered_number, pin) VALUES ($1, $2, $3);

-- name: CreateWallet :exec
INSERT INTO wallets (id, balance, last_update) VALUES ($1, $2, DEFAULT);

-- name: CreateAccount :exec
INSERT INTO accounts (id, users, auth_info, wallet, level) VALUES ($1, $2, $3, $4, $5);

-- name: CreateTransfer :exec
INSERT INTO transfers (id, from_wallet, to_wallet, balance, transfer_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: GetUserWallet :one
SELECT a.wallet FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1;

-- name: GetUserWalletAndAuth :one
SELECT a.wallet, ai.pin FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1;