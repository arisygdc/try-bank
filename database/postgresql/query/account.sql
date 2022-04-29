-- name: AccountType :one
SELECT * FROM account_type WHERE name = $1;

-- name: CreateCustomer :exec
INSERT INTO cutomers (id, firstname, lastname, created_at, email, birth, phone) VALUES ($1, $2, $3, DEFAULT, $4, $5, $6);

-- name: CreateAuthInfo :one
INSERT INTO auth_info (id, registered_number, pin) VALUES ($1, DEFAULT, $2) RETURNING registered_number;

-- name: CreateWallet :exec
INSERT INTO wallets (id, balance, last_update) VALUES ($1, $2, DEFAULT);

-- name: CreateAccount :exec
INSERT INTO accounts (id, cutomer_id, auth_info_id, wallet_id, account_type_id) VALUES ($1, $2, $3, $4, $5);

-- name: CreateTransfer :exec
INSERT INTO transfers (id, from_wallet, to_wallet, balance, transfer_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: GetUserWallet :one
SELECT a.wallet_id FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1;

-- name: GetUserWalletAndAuth :one
SELECT a.wallet_id, ai.pin FROM accounts a
RIGHT JOIN auth_info ai ON ai.id = a.auth_info
WHERE ai.registered_number = $1;