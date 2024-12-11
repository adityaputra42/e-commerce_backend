// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: orders.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
 id,
 transaction_id,
 product_id,
 color_varian_id,
 size_varian_id,
 unit_price,
 subtotal,
 quantity,
 status
) VALUES (
  $1, $2 ,$3, $4, $5, $6, $7, $8, $9
)
RETURNING id, transaction_id, product_id, color_varian_id, size_varian_id, unit_price, subtotal, quantity, status, updated_at, created_at, deleted_at
`

type CreateOrderParams struct {
	ID            string  `json:"id"`
	TransactionID string  `json:"transaction_id"`
	ProductID     int64   `json:"product_id"`
	ColorVarianID int64   `json:"color_varian_id"`
	SizeVarianID  int64   `json:"size_varian_id"`
	UnitPrice     float64 `json:"unit_price"`
	Subtotal      float64 `json:"subtotal"`
	Quantity      int64   `json:"quantity"`
	Status        string  `json:"status"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder,
		arg.ID,
		arg.TransactionID,
		arg.ProductID,
		arg.ColorVarianID,
		arg.SizeVarianID,
		arg.UnitPrice,
		arg.Subtotal,
		arg.Quantity,
		arg.Status,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.TransactionID,
		&i.ProductID,
		&i.ColorVarianID,
		&i.SizeVarianID,
		&i.UnitPrice,
		&i.Subtotal,
		&i.Quantity,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
UPDATE orders
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteOrder, id)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT 
o.id AS id, 
o.transaction_id AS transaction_id,
jsonb_build_object(
        'id', p.id,
        'name', p.name,
        'category',(
             SELECT jsonb_build_object(
                    'id', c.id,
                    'name', c.name,
                    'icon', c.icon
                   ) FROM categories c
              WHERE c.id = p.category_id AND c.deleted_at IS NOT NULL
        ),
        'description', p.description,
        'images',p.images,
        'rating',p.rating,
        'price', p.price,
        'created_at',p.created_at,
        'updated_at',p.updated_at,
        'color_varian',(
               SELECT jsonb_build_object(
                    'id', cv.id,
                    'name',cv.name,
                    'color',cv.color,
                    'images',cv.images,
                    'created_at',cv.created_at,
                    'updated_at',cv.updated_at
               ) FROM color_varians cv
          WHERE cv.id = o.color_varian_id AND cv.deleted_at IS NOT NULL 
        )
    ) AS product, ( 
     SELECT jsonb_build_object(
      'id', a.id,
      'recipient_name', a.recipient_name,
      'recipient_phone_number', a.recipient_phone_number,
      'province',a.province,
      'city',a.city,
      'district',a.district,
      'village',a.village,
      'postal_code',a.postal_code,
      'full_address',a.full_address     
     ) FROM address a WHERE a.id = tx.address_id
    ) AS address,
    sv.size AS size,
    o.subtotal AS subtotal,
    o.quantity AS quantity,
    o.status AS status,
    o.created_at AS created_at,
    o.updated_at AS updated_at
     FROM orders o LEFT JOIN 
     products p ON o.product_id = p.id AND p.deleted_at IS NOT NULL
     LEFT JOIN size_varians sv ON o.size_varian_id = sv.id AND sv.deleted_at IS NOT NULL
     LEFT JOIN transactions tx ON o.transaction_id = tx.tx_id AND tx.deleted_at IS NOT NULL
     WHERE o.id = $1 LIMIT 1 AND o.deleted_at IS NOT NULL
`

type GetOrderRow struct {
	ID            string      `json:"id"`
	TransactionID string      `json:"transaction_id"`
	Product       []byte      `json:"product"`
	Address       []byte      `json:"address"`
	Size          pgtype.Text `json:"size"`
	Subtotal      float64     `json:"subtotal"`
	Quantity      int64       `json:"quantity"`
	Status        string      `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

func (q *Queries) GetOrder(ctx context.Context, id string) (GetOrderRow, error) {
	row := q.db.QueryRow(ctx, getOrder, id)
	var i GetOrderRow
	err := row.Scan(
		&i.ID,
		&i.TransactionID,
		&i.Product,
		&i.Address,
		&i.Size,
		&i.Subtotal,
		&i.Quantity,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrderForUpdate = `-- name: GetOrderForUpdate :one
SELECT id, transaction_id, product_id, color_varian_id, size_varian_id, unit_price, subtotal, quantity, status, updated_at, created_at, deleted_at FROM orders
WHERE deleted_at IS NOT NULL AND id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetOrderForUpdate(ctx context.Context, id string) (Order, error) {
	row := q.db.QueryRow(ctx, getOrderForUpdate, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.TransactionID,
		&i.ProductID,
		&i.ColorVarianID,
		&i.SizeVarianID,
		&i.UnitPrice,
		&i.Subtotal,
		&i.Quantity,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listOrder = `-- name: ListOrder :many
SELECT 
    o.id AS id, 
    o.transaction_id AS transaction_id,
    jsonb_build_object(
        'id', p.id,
        'name', p.name,
        'category',(
             SELECT jsonb_build_object(
                    'id', c.id,
                    'name', c.name,
                    'icon', c.icon
                   ) FROM categories c
              WHERE c.id = p.category_id AND c.deleted_at IS NOT NULL
        ),
        'description', p.description,
        'images',p.images,
        'rating',p.rating,
        'price', p.price,
        'created_at',p.created_at,
        'updated_at',p.updated_at,
        'color_varian',(
               SELECT jsonb_build_object(
                    'id', cv.id,
                    'name',cv.name,
                    'color',cv.color,
                    'images',cv.images,
                    'created_at',cv.created_at,
                    'updated_at',cv.updated_at
               ) FROM color_varians cv
          WHERE cv.id = o.color_varian_id AND cv.deleted_at IS NOT NULL
        )
    ) AS product,
    sv.size AS size,
    o.subtotal AS subtotal,
    o.quantity AS quantity,
    o.status AS status,
    o.created_at AS created_at,
    o.updated_at AS updated_at
 FROM orders o 
 LEFT JOIN products p ON o.product_id = p.id AND p.deleted_at IS NOT NULL
 LEFT JOIN size_varians sv ON o.size_varian_id = sv.id AND sv.deleted_at IS NOT NULL
 WHERE o.deleted_at IS NOT NULL
ORDER BY o.id
LIMIT $1
OFFSET $2
`

type ListOrderParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListOrderRow struct {
	ID            string      `json:"id"`
	TransactionID string      `json:"transaction_id"`
	Product       []byte      `json:"product"`
	Size          pgtype.Text `json:"size"`
	Subtotal      float64     `json:"subtotal"`
	Quantity      int64       `json:"quantity"`
	Status        string      `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

func (q *Queries) ListOrder(ctx context.Context, arg ListOrderParams) ([]ListOrderRow, error) {
	rows, err := q.db.Query(ctx, listOrder, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListOrderRow{}
	for rows.Next() {
		var i ListOrderRow
		if err := rows.Scan(
			&i.ID,
			&i.TransactionID,
			&i.Product,
			&i.Size,
			&i.Subtotal,
			&i.Quantity,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateOrder = `-- name: UpdateOrder :one
UPDATE orders
 set status = $2
WHERE id = $1
RETURNING id, transaction_id, product_id, color_varian_id, size_varian_id, unit_price, subtotal, quantity, status, updated_at, created_at, deleted_at
`

type UpdateOrderParams struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, updateOrder, arg.ID, arg.Status)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.TransactionID,
		&i.ProductID,
		&i.ColorVarianID,
		&i.SizeVarianID,
		&i.UnitPrice,
		&i.Subtotal,
		&i.Quantity,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
