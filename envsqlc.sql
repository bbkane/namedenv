-- name: CreateEnv :one
INSERT INTO env (
    name, comment, create_time, update_time
) VALUES (
    ?   , ?      , ?          , ?
)
RETURNING id;

-- See https://github.com/sqlc-dev/sqlc/issues/937#issuecomment-798858187 for why NULLIF is needed
-- name: UpdateEnv :exec
UPDATE env SET
    name = COALESCE(sqlc.narg('name'), name),
    comment = COALESCE(sqlc.narg('comment'), comment),
    create_time = COALESCE(sqlc.narg('create_time'), create_time),
    update_time = COALESCE(sqlc.narg('update_time'), update_time)
WHERE id = sqlc.arg('id');