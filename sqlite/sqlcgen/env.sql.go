// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: env.sql

package sqlcgen

import (
	"context"
)

const createEnv = `-- name: CreateEnv :one
INSERT INTO env (
    name, comment, create_time, update_time
) VALUES (
    ?   , ?      , ?          , ?
)
RETURNING name, comment, create_time, update_time
`

type CreateEnvParams struct {
	Name       string
	Comment    *string
	CreateTime string
	UpdateTime string
}

type CreateEnvRow struct {
	Name       string
	Comment    *string
	CreateTime string
	UpdateTime string
}

func (q *Queries) CreateEnv(ctx context.Context, arg CreateEnvParams) (CreateEnvRow, error) {
	row := q.db.QueryRowContext(ctx, createEnv,
		arg.Name,
		arg.Comment,
		arg.CreateTime,
		arg.UpdateTime,
	)
	var i CreateEnvRow
	err := row.Scan(
		&i.Name,
		&i.Comment,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const findEnv = `-- name: FindEnv :one
SELECT
    name, comment, create_time, update_time
FROM env
WHERE name = ?
`

type FindEnvRow struct {
	Name       string
	Comment    *string
	CreateTime string
	UpdateTime string
}

func (q *Queries) FindEnv(ctx context.Context, name string) (FindEnvRow, error) {
	row := q.db.QueryRowContext(ctx, findEnv, name)
	var i FindEnvRow
	err := row.Scan(
		&i.Name,
		&i.Comment,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const findEnvID = `-- name: FindEnvID :one
SELECT id FROM env WHERE name = ?
`

func (q *Queries) FindEnvID(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, findEnvID, name)
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
WHERE name = ?5
`

type UpdateEnvParams struct {
	NewName    *string
	Comment    *string
	CreateTime *string
	UpdateTime *string
	Name       string
}

// See https://docs.sqlc.dev/en/latest/howto/named_parameters.html#nullable-parameters
func (q *Queries) UpdateEnv(ctx context.Context, arg UpdateEnvParams) error {
	_, err := q.db.ExecContext(ctx, updateEnv,
		arg.NewName,
		arg.Comment,
		arg.CreateTime,
		arg.UpdateTime,
		arg.Name,
	)
	return err
}
