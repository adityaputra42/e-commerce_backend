// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: size_varian.sql

package db

import (
	"context"
)

const createSizeVarianProduct = `-- name: CreateSizeVarianProduct :one
INSERT INTO size_varians (
 color_varian_id,
 size,
 stock
) VALUES (
  $1, $2 ,$3
)
RETURNING id, color_varian_id, size, stock, updated_at, created_at, deleted_at
`

type CreateSizeVarianProductParams struct {
	ColorVarianID int64  `json:"color_varian_id"`
	Size          string `json:"size"`
	Stock         int64  `json:"stock"`
}

func (q *Queries) CreateSizeVarianProduct(ctx context.Context, arg CreateSizeVarianProductParams) (SizeVarian, error) {
	row := q.db.QueryRow(ctx, createSizeVarianProduct, arg.ColorVarianID, arg.Size, arg.Stock)
	var i SizeVarian
	err := row.Scan(
		&i.ID,
		&i.ColorVarianID,
		&i.Size,
		&i.Stock,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteSizeVarianProduct = `-- name: DeleteSizeVarianProduct :exec
UPDATE size_varians
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) DeleteSizeVarianProduct(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSizeVarianProduct, id)
	return err
}

const getSizeVarianProduct = `-- name: GetSizeVarianProduct :one
SELECT id, color_varian_id, size, stock, updated_at, created_at, deleted_at FROM size_varians
WHERE deleted_at IS NOT NULL AND id = $1 LIMIT 1
`

func (q *Queries) GetSizeVarianProduct(ctx context.Context, id int64) (SizeVarian, error) {
	row := q.db.QueryRow(ctx, getSizeVarianProduct, id)
	var i SizeVarian
	err := row.Scan(
		&i.ID,
		&i.ColorVarianID,
		&i.Size,
		&i.Stock,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getSizeVarianProductForUpdate = `-- name: GetSizeVarianProductForUpdate :one
SELECT id, color_varian_id, size, stock, updated_at, created_at, deleted_at FROM size_varians
WHERE deleted_at IS NOT NULL AND id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetSizeVarianProductForUpdate(ctx context.Context, id int64) (SizeVarian, error) {
	row := q.db.QueryRow(ctx, getSizeVarianProductForUpdate, id)
	var i SizeVarian
	err := row.Scan(
		&i.ID,
		&i.ColorVarianID,
		&i.Size,
		&i.Stock,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listSizeVarianProduct = `-- name: ListSizeVarianProduct :many
SELECT id, color_varian_id, size, stock, updated_at, created_at, deleted_at FROM size_varians
WHERE deleted_at IS NOT NULL AND color_varian_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListSizeVarianProductParams struct {
	ColorVarianID int64 `json:"color_varian_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListSizeVarianProduct(ctx context.Context, arg ListSizeVarianProductParams) ([]SizeVarian, error) {
	rows, err := q.db.Query(ctx, listSizeVarianProduct, arg.ColorVarianID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SizeVarian{}
	for rows.Next() {
		var i SizeVarian
		if err := rows.Scan(
			&i.ID,
			&i.ColorVarianID,
			&i.Size,
			&i.Stock,
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

const updateSizeVarianProduct = `-- name: UpdateSizeVarianProduct :one
UPDATE size_varians
 set size = $2,
 stock = $3
WHERE id = $1
RETURNING id, color_varian_id, size, stock, updated_at, created_at, deleted_at
`

type UpdateSizeVarianProductParams struct {
	ID    int64  `json:"id"`
	Size  string `json:"size"`
	Stock int64  `json:"stock"`
}

func (q *Queries) UpdateSizeVarianProduct(ctx context.Context, arg UpdateSizeVarianProductParams) (SizeVarian, error) {
	row := q.db.QueryRow(ctx, updateSizeVarianProduct, arg.ID, arg.Size, arg.Stock)
	var i SizeVarian
	err := row.Scan(
		&i.ID,
		&i.ColorVarianID,
		&i.Size,
		&i.Stock,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
