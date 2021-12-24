-- name: CreateCompany :exec
INSERT INTO companies (id, name, email, phone, created_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: CreateCompanyAccount :exec
INSERT INTO companies_account (id, company, auth_info, wallet, virtual_account) VALUES ($1, $2, $3, $4, $5);

-- name: CreateVirtualAccount :exec
INSERT INTO virtual_account (id, va_key, domain, va_identity, created_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: CreateVAPayment :exec
INSERT INTO va_payment (id, virtual_account, va_number, request_payment, paid_at) VALUES (NULL, $1, $2, $3, DEFAULT);

-- name: UpdateVAstatus :exec
UPDATE companies_account SET virtual_account = $1 WHERE id = $2;

-- name: ValidateCompany :one
SELECT ca.id FROM companies_account ca 
RIGHT JOIN companies c ON ca.company = c.id 
RIGHT JOIN auth_info ai ON ca.auth_info = ai.id
WHERE c.name = $1 AND c.email = $2 AND c.phone = $3
AND ai.registered_number = $4;