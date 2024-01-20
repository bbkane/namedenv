-- name: CreateEnv :one
INSERT INTO env (
    name, comment, create_time, update_time
) VALUES (
    ?   , ?      , ?          , ?
)
RETURNING name, comment, create_time, update_time;

-- name: DeleteEnv :exec
DELETE FROM env WHERE name = ?;

-- name: FindEnvID :one
SELECT id FROM env WHERE name = ?;

-- name: FindEnv :one
SELECT
    name, comment, create_time, update_time
FROM env
WHERE name = ?;

-- See https://docs.sqlc.dev/en/latest/howto/named_parameters.html#nullable-parameters
-- name: UpdateEnv :exec
UPDATE env SET
    name = COALESCE(sqlc.narg('new_name'), name),
    comment = COALESCE(sqlc.narg('comment'), comment),
    create_time = COALESCE(sqlc.narg('create_time'), create_time),
    update_time = COALESCE(sqlc.narg('update_time'), update_time)
WHERE name = sqlc.arg('name');

