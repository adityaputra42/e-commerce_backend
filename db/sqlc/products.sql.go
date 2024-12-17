// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package db

import (
	"context"
	"time"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
 category_id,
 name,
 description,
 images,
 rating,
 price
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, category_id, name, description, images, rating, price, updated_at, created_at, deleted_at
`

type CreateProductParams struct {
	CategoryID  int64   `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Images      string  `json:"images"`
	Rating      float64 `json:"rating"`
	Price       float64 `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.CategoryID,
		arg.Name,
		arg.Description,
		arg.Images,
		arg.Rating,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Images,
		&i.Rating,
		&i.Price,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, category_id, name, description, images, rating, price, updated_at, created_at, deleted_at FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Images,
		&i.Rating,
		&i.Price,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProductForUpdate = `-- name: GetProductForUpdate :one
SELECT id, category_id, name, description, images, rating, price, updated_at, created_at, deleted_at FROM products
WHERE deleted_at IS NOT NULL AND id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetProductForUpdate(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProductForUpdate, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Images,
		&i.Rating,
		&i.Price,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProductWithDetail = `-- name: GetProductWithDetail :one
SELECT 
    p.id AS product_id,
    jsonb_build_object(
        'id', c.id,
        'name', c.name,
        'icon', c.icon
    ) AS category,
    p.name AS name,
    p.description AS description,
    p.images AS images,
    p.rating AS rating,
    p.price AS price,
    p.updated_at AS updated_at,
    p.created_at AS created_at,
    (
        SELECT jsonb_agg(
            jsonb_build_object(
                'id', cv.id,
                'product_id', cv.product_id,
                'name', cv.name,
                'color', cv.color,
                'images', cv.images,
                'updated_at', cv.updated_at,
                'created_at', cv.created_at,
                'size_varian', (
                    SELECT jsonb_agg(
                        jsonb_build_object(
                            'id', sv.id,
                            'color_varian_id', sv.color_varian_id,
                            'size', sv.size,
                            'stock', sv.stock,
                            'updated_at', sv.updated_at,
                            'created_at', sv.created_at
                        )
                    )
                    FROM size_varians sv
                    WHERE sv.color_varian_id = cv.id AND sv.deleted_at IS NOT NULL
                )
            )
        )
        FROM color_varians cv
        WHERE cv.product_id = p.id AND cv.deleted_at IS NOT NULL
    ) AS color_varian
FROM 
    products p
LEFT JOIN 
    categories c ON p.category_id = c.id AND c.deleted_at IS NOT NULL
WHERE 
p.deleted_at IS NOT NULL AND
p.id = $1
`

type GetProductWithDetailRow struct {
	ProductID   int64     `json:"product_id"`
	Category    []byte    `json:"category"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Images      string    `json:"images"`
	Rating      float64   `json:"rating"`
	Price       float64   `json:"price"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	ColorVarian []byte    `json:"color_varian"`
}

func (q *Queries) GetProductWithDetail(ctx context.Context, id int64) (GetProductWithDetailRow, error) {
	row := q.db.QueryRow(ctx, getProductWithDetail, id)
	var i GetProductWithDetailRow
	err := row.Scan(
		&i.ProductID,
		&i.Category,
		&i.Name,
		&i.Description,
		&i.Images,
		&i.Rating,
		&i.Price,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.ColorVarian,
	)
	return i, err
}

const listProduct = `-- name: ListProduct :many
SELECT 
    p.id AS product_id,
    jsonb_build_object(
        'id', c.id,
        'name', c.name,
        'icon', c.icon
    ) AS category,
    p.name AS name,
    p.description AS description,
    p.images AS images,
    p.rating AS rating,
    p.price AS price,
    p.updated_at AS updated_at,
    p.created_at AS created_at 
FROM 
    products p
JOIN 
    categories c ON p.category_id = c.id AND c.deleted_at IS NOT NULL
WHERE p.deleted_at IS NOT NULL
ORDER BY p.id
LIMIT $1
OFFSET $2
`

type ListProductParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListProductRow struct {
	ProductID   int64     `json:"product_id"`
	Category    []byte    `json:"category"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Images      string    `json:"images"`
	Rating      float64   `json:"rating"`
	Price       float64   `json:"price"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func (q *Queries) ListProduct(ctx context.Context, arg ListProductParams) ([]ListProductRow, error) {
	rows, err := q.db.Query(ctx, listProduct, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductRow{}
	for rows.Next() {
		var i ListProductRow
		if err := rows.Scan(
			&i.ProductID,
			&i.Category,
			&i.Name,
			&i.Description,
			&i.Images,
			&i.Rating,
			&i.Price,
			&i.UpdatedAt,
			&i.CreatedAt,
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

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
 set category_id = $2,
 name = $3,
 description= $4,
 images= $5,
 rating= $6,
 price= $7
WHERE id = $1
RETURNING id, category_id, name, description, images, rating, price, updated_at, created_at, deleted_at
`

type UpdateProductParams struct {
	ID          int64   `json:"id"`
	CategoryID  int64   `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Images      string  `json:"images"`
	Rating      float64 `json:"rating"`
	Price       float64 `json:"price"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.CategoryID,
		arg.Name,
		arg.Description,
		arg.Images,
		arg.Rating,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Description,
		&i.Images,
		&i.Rating,
		&i.Price,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
