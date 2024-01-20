-- name: EnvLocalVarCreate :exec
INSERT INTO env_var_local(
    env_id, name, comment, create_time, update_time, value
) VALUES (
    ?     , ?   , ?      , ?          , ?          , ?
);

-- name: EnvLocalVarList :many
SELECT * FROM env_var_local
WHERE env_id = ?
ORDER BY name ASC;

-- name: FindEnvLocalVar :one
SELECT *
FROM env_var_local
WHERE env_id = ? AND name = ?;