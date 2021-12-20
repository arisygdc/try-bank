-- name: CreatePermissionLevel :exec
INSERT INTO permission_level (id, name) VALUES ($1, $2);

-- name: CreateUser :exec
INSERT INTO users (id, firstname, lastname, created_at, email, birth, phone) VALUES ($1, $2, $3, DEFAULT, $4, $5, $6);

-- name: CreateAuthInfo :exec
INSERT INTO auth_info (id, registered_number, pin) VALUES ($1, $2, $3);

-- name: CreateCoustomerWallet :exec
INSERT INTO coustomer_wallet (id, balance, last_update) VALUES ($1, $2, DEFAULT);

-- name: CreateAccount :exec
INSERT INTO accounts (id, users, auth_info, wallet, permission) VALUES ($1, $2, $3, $4, $5);

-- name: CreateCompany :exec
INSERT INTO companies (id, name, company_key) VALUES ($1, $2, $3);

-- name: CreateCompanyWallet :exec
INSERT INTO companies_wallet (id, balance) VALUES ($1, $2);

-- name: CreateAccountHaveCompany :exec
INSERT INTO account_have_company (account, company, company_wallet) VALUES ($1, $2, $3);

-- name: CreateVirtualAccount :exec
INSERT INTO virtual_account (id, company_id, request_payment, va_number, paid_at) VALUES (NULL, $1, $2, $3, DEFAULT);

-- name: CreateTransfer :exec
INSERT INTO transfers (id, from_account, to_account, balance, transfer_at) VALUES ($1, $2, $3, $3, DEFAULT);