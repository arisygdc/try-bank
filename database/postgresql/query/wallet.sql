-- name: SubtractBalance :execrows
UPDATE wallets SET balance = (balance - $1) WHERE id = $2;

-- name: AddBalance :execrows
UPDATE wallets SET balance = (balance + $1) WHERE id = $2;