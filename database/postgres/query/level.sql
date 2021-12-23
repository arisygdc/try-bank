-- name: CreatePermissionLevel :exec
INSERT INTO levels (id, name) VALUES ($1, $2);

-- name: GetPermissionID :one
SELECT id FROM levels WHERE name = $1;