-- name: CreateCompany :exec
INSERT INTO companies (id, name, email, phone, created_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: CreateCompanyAccount :exec
INSERT INTO companies_account (id, company_id, auth_info_id, wallet_id, virtual_account_id) VALUES ($1, $2, $3, $4, $5);

-- name: CreateVirtualAccount :exec
INSERT INTO virtual_account (id, authorization_key, identity, callback_url, created_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: SetCompanyVA :exec
UPDATE companies_account SET virtual_account_id = $1 WHERE id = $2;

-- name: CreateVAPayment :exec
INSERT INTO va_payment (id, issued_payment_id, paid_at) VALUES ($1, $2, DEFAULT);

-- name: ValidateCompany :one
SELECT ca.id FROM companies_account ca 
RIGHT JOIN companies c ON ca.company = c.id 
RIGHT JOIN auth_info ai ON ca.auth_info = ai.id
WHERE c.name = $1 AND c.email = $2 AND c.phone = $3
AND ai.registered_number = $4;

-- name: AuthGetCompaniesAccount :one
SELECT company_id, auth_info_id, wallet_id, virtual_account_id
FROM companies_account ca 
LEFT JOIN auth_info ai 
ON ai.id = ca.auth_info_id 
WHERE ai.registered_number = $1;