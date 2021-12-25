-- name: CheckVA :one
SELECT ca.id as comp_id, va.id as va_id, ca.wallet as wallet_id, va.va_key, va.fqdn_detail_payment, fqdn_pay 
FROM virtual_account va
LEFT JOIN companies_account ca ON va.id = ca.virtual_account
WHERE va_identity = $1;

-- name: PayVA :exec
INSERT INTO va_payment (id, virtual_account, va_number, request_payment, paid_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: AddBalance :exec
UPDATE wallets SET balance = balance + $1 WHERE id = $2;

-- name: SubtractBalance :exec
UPDATE wallets SET balance = balance - $1 WHERE id = $2;

-- name: GetBalance :one
SELECT balance FROM wallets WHERE id = $1;