-- name: CreateLevel :exec
INSERT INTO levels (id, name) VALUES ($1, $2);

-- name: GetLevelID :one
SELECT id FROM levels WHERE name = $1;