// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: env_var_local.sql

package sqlcgen

import (
	"context"
)

const envVarCreate = `-- name: EnvVarCreate :exec
INSERT INTO env_var_local(
    env_id, name, comment, create_time, update_time, value
) VALUES (
    ?     , ?   , ?      , ?          , ?          , ?
)
`

type EnvVarCreateParams struct {
	EnvID      int64
	Name       string
	Comment    string
	CreateTime string
	UpdateTime string
	Value      string
}

func (q *Queries) EnvVarCreate(ctx context.Context, arg EnvVarCreateParams) error {
	_, err := q.db.ExecContext(ctx, envVarCreate,
		arg.EnvID,
		arg.Name,
		arg.Comment,
		arg.CreateTime,
		arg.UpdateTime,
		arg.Value,
	)
	return err
}

const envVarDelete = `-- name: EnvVarDelete :exec
DELETE FROM env_var_local WHERE env_id = ? AND name = ?
`

type EnvVarDeleteParams struct {
	EnvID int64
	Name  string
}

func (q *Queries) EnvVarDelete(ctx context.Context, arg EnvVarDeleteParams) error {
	_, err := q.db.ExecContext(ctx, envVarDelete, arg.EnvID, arg.Name)
	return err
}

const envVarFindByID = `-- name: EnvVarFindByID :one
SELECT id, env_id, name, comment, create_time, update_time, value FROM env_var_local WHERE id = ?
`

func (q *Queries) EnvVarFindByID(ctx context.Context, id int64) (EnvVarLocal, error) {
	row := q.db.QueryRowContext(ctx, envVarFindByID, id)
	var i EnvVarLocal
	err := row.Scan(
		&i.ID,
		&i.EnvID,
		&i.Name,
		&i.Comment,
		&i.CreateTime,
		&i.UpdateTime,
		&i.Value,
	)
	return i, err
}

const envVarFindID = `-- name: EnvVarFindID :one
SELECT id FROM env_var_local WHERE env_id = ? AND name = ?
`

type EnvVarFindIDParams struct {
	EnvID int64
	Name  string
}

func (q *Queries) EnvVarFindID(ctx context.Context, arg EnvVarFindIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, envVarFindID, arg.EnvID, arg.Name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const envVarList = `-- name: EnvVarList :many
SELECT id, env_id, name, comment, create_time, update_time, value FROM env_var_local
WHERE env_id = ?
ORDER BY name ASC
`

func (q *Queries) EnvVarList(ctx context.Context, envID int64) ([]EnvVarLocal, error) {
	rows, err := q.db.QueryContext(ctx, envVarList, envID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EnvVarLocal
	for rows.Next() {
		var i EnvVarLocal
		if err := rows.Scan(
			&i.ID,
			&i.EnvID,
			&i.Name,
			&i.Comment,
			&i.CreateTime,
			&i.UpdateTime,
			&i.Value,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const envVarShow = `-- name: EnvVarShow :one
SELECT id, env_id, name, comment, create_time, update_time, value
FROM env_var_local
WHERE env_id = ? AND name = ?
`

type EnvVarShowParams struct {
	EnvID int64
	Name  string
}

func (q *Queries) EnvVarShow(ctx context.Context, arg EnvVarShowParams) (EnvVarLocal, error) {
	row := q.db.QueryRowContext(ctx, envVarShow, arg.EnvID, arg.Name)
	var i EnvVarLocal
	err := row.Scan(
		&i.ID,
		&i.EnvID,
		&i.Name,
		&i.Comment,
		&i.CreateTime,
		&i.UpdateTime,
		&i.Value,
	)
	return i, err
}
