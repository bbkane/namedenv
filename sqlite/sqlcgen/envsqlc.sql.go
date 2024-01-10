// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: envsqlc.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const createEnv = `-- name: CreateEnv :one
INSERT INTO env (
    name, comment, create_time, update_time
) VALUES (
    ?   , ?      , ?          , ?
)
RETURNING id
`

type CreateEnvParams struct {
	Name       string
	Comment    sql.NullString
	CreateTime string
	UpdateTime string
}

func (q *Queries) CreateEnv(ctx context.Context, arg CreateEnvParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createEnv,
		arg.Name,
		arg.Comment,
		arg.CreateTime,
		arg.UpdateTime,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateEnv = `-- name: UpdateEnv :exec
UPDATE env SET
    name = COALESCE(?1, name),
    comment = COALESCE(?2, comment),
    create_time = COALESCE(?3, create_time),
    update_time = COALESCE(?4, update_time)
WHERE id = ?5
`

type UpdateEnvParams struct {
	Name       sql.NullString
	Comment    sql.NullString
	CreateTime sql.NullString
	UpdateTime sql.NullString
	ID         int64
}

// See https://github.com/sqlc-dev/sqlc/issues/937#issuecomment-798858187 for why NULLIF is needed
func (q *Queries) UpdateEnv(ctx context.Context, arg UpdateEnvParams) error {
	_, err := q.db.ExecContext(ctx, updateEnv,
		arg.Name,
		arg.Comment,
		arg.CreateTime,
		arg.UpdateTime,
		arg.ID,
	)
	return err
}
