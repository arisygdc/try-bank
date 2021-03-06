-- name: CreateVirtualAccount :exec
INSERT INTO virtual_account (id, authorization_key, identity, callback_url, created_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: ActivateVirtualAccount :execrows
UPDATE companies_account SET virtual_account_id = $1 WHERE company_id = $2;

-- name: IssuePaymentVA :exec
INSERT INTO issued_payment (id, virtual_account_id, virtual_account_number, payment_charge) VALUES ($1, $2, $3, $4);

-- name: VirtualAccountID :one
SELECT id FROM virtual_account WHERE identity = $1;

-- name: PaymentVA :one
INSERT INTO va_payment (id, issued_payment_id) VALUES ($1, $2) RETURNING *;

-- name: CheckActiveIssueVAP :one
SELECT * FROM issued_payment 
WHERE virtual_account_id = $1
AND virtual_account_number = $2
AND issued_at + INTERVAL '1 day' > NOW()
AND id NOT IN (SELECT issued_payment_id FROM va_payment);