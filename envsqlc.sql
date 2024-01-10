-- name: CreateEnv :one
INSERT INTO env (
    name, comment, create_time, update_time
) VALUES (
    ?   , ?      , ?          , ?
)
RETURNING id;

-- name: FindEnvByName :one
SELECT id FROM env WHERE name = ?;

-- See https://docs.sqlc.dev/en/latest/howto/named_parameters.html#nullable-parameters
-- name: UpdateEnv :exec
UPDATE env SET
    name = COALESCE(sqlc.narg('name'), name),
    comment = COALESCE(sqlc.narg('comment'), comment),
    create_time = COALESCE(sqlc.narg('create_time'), create_time),
    update_time = COALESCE(sqlc.narg('update_time'), update_time)
WHERE id = sqlc.arg('id');