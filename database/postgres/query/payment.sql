-- name: CheckVA :one
SELECT ca.id as comp_id, va.id as va_id, ca.wallet as wallet_id, va.va_key, va.domain FROM virtual_account va
LEFT JOIN companies_account ca ON va.id = ca.virtual_account
WHERE va_identity = $1;

-- name: PayVA :exec
INSERT INTO va_payment (id, virtual_account, va_number, request_payment, paid_at) VALUES ($1, $2, $3, $4, DEFAULT);

-- name: UpdateBalance :exec
UPDATE wallets SET balance = balance + $1 WHERE id = $2;