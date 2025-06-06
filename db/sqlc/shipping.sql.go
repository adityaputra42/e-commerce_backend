// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: shipping.sql

package db

import (
	"context"
)

const createShipping = `-- name: CreateShipping :one
INSERT INTO shippings (
 name,
 price,
 state
) VALUES (
  $1, $2 ,$3
)
RETURNING id, name, price, state, updated_at, created_at, deleted_at
`

type CreateShippingParams struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	State string  `json:"state"`
}

func (q *Queries) CreateShipping(ctx context.Context, arg CreateShippingParams) (Shipping, error) {
	row := q.db.QueryRow(ctx, createShipping, arg.Name, arg.Price, arg.State)
	var i Shipping
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.State,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteShipping = `-- name: DeleteShipping :exec
UPDATE shippings
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) DeleteShipping(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteShipping, id)
	return err
}

const getShipping = `-- name: GetShipping :one
SELECT id, name, price, state, updated_at, created_at, deleted_at FROM shippings
WHERE deleted_at IS NULL AND id = $1 LIMIT 1
`

func (q *Queries) GetShipping(ctx context.Context, id int64) (Shipping, error) {
	row := q.db.QueryRow(ctx, getShipping, id)
	var i Shipping
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.State,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getShippingForUpdate = `-- name: GetShippingForUpdate :one
SELECT id, name, price, state, updated_at, created_at, deleted_at FROM shippings
WHERE deleted_at IS NULL AND id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetShippingForUpdate(ctx context.Context, id int64) (Shipping, error) {
	row := q.db.QueryRow(ctx, getShippingForUpdate, id)
	var i Shipping
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.State,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listShipping = `-- name: ListShipping :many
SELECT id, name, price, state, updated_at, created_at, deleted_at FROM shippings
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListShippingParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListShipping(ctx context.Context, arg ListShippingParams) ([]Shipping, error) {
	rows, err := q.db.Query(ctx, listShipping, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Shipping{}
	for rows.Next() {
		var i Shipping
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.State,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateShipping = `-- name: UpdateShipping :one
UPDATE shippings
 set name = $2,
 price = $3,
 state = $4
WHERE id = $1
RETURNING id, name, price, state, updated_at, created_at, deleted_at
`

type UpdateShippingParams struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	State string  `json:"state"`
}

func (q *Queries) UpdateShipping(ctx context.Context, arg UpdateShippingParams) (Shipping, error) {
	row := q.db.QueryRow(ctx, updateShipping,
		arg.ID,
		arg.Name,
		arg.Price,
		arg.State,
	)
	var i Shipping
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.State,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
