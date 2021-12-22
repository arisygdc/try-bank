-- name: CreatePermissionLevel :exec
INSERT INTO permission_level (id, name) VALUES ($1, $2);

-- name: GetPermissionID :one
SELECT id FROM permission_level WHERE name = $1;