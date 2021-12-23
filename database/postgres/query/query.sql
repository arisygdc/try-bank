-- name: CreateUser :exec
INSERT INTO users (id, firstname, lastname, created_at, email, birth, phone) VALUES ($1, $2, $3, DEFAULT, $4, $5, $6);

-- name: CreateAuthInfo :exec
INSERT INTO auth_info (id, registered_number, pin) VALUES ($1, $2, $3);

-- name: CreateWallet :exec
INSERT INTO wallets (id, balance, last_update) VALUES ($1, $2, DEFAULT);

-- name: CreateAccount :exec
INSERT INTO accounts (id, users, auth_info, wallet, level) VALUES ($1, $2, $3, $4, $5);

-- name: CreateCompany :exec
INSERT INTO companies (id, name, email, created_at) VALUES ($1, $2, $3, DEFAULT);

-- name: CreateCompanyAccount :exec
INSERT INTO companies_account (id, company, auth_info, wallet, virtual_account) VALUES ($1, $2, $3, $4, $5);

-- name: CreateVirtualAccount :exec
INSERT INTO virtual_account (id, va_key, domain, va_identity, created_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: CreateVAPayment :exec
INSERT INTO va_payment (id, virtual_account, va_number, request_payment, paid_at) VALUES (NULL, $1, $2, $3, DEFAULT);

-- name: CreateTransfer :exec
INSERT INTO transfers (id, from_wallet, to_wallet, balance, transfer_at) VALUES ($1, $2, $3, $3, DEFAULT);