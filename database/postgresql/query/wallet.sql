-- name: SubtractBalance :execrows
UPDATE wallets SET balance = (balance - $1) WHERE id = $2;

-- name: AddBalance :execrows
UPDATE wallets SET balance = (balance + $1) WHERE id = $2;

-- name: Transfer :exec
INSERT INTO transfers (id, from_wallet, to_wallet, balance, transfered_at) VALUES ($1, $2, $3, $4, DEFAULT);