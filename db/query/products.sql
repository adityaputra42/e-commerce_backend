-- name: CreateProduct :one
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
RETURNING *;


-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductWithDetail :one
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
                    WHERE sv.color_varian_id = cv.id
                )
            )
        )
        FROM color_varians cv
        WHERE cv.product_id = p.id
    ) AS color_varian
FROM 
    products p
JOIN 
    categories c ON p.category_id = c.id
WHERE 
    p.id = $1;

-- name: GetProductForUpdate :one
SELECT * FROM products
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListProduct :many
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
    categories c ON p.category_id = c.id
ORDER BY p.id
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
 set category_id = $2,
 name = $3,
 description= $4,
 images= $5,
 rating= $6,
 price= $7
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;