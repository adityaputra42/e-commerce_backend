-- name: CreateOrder :one
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
RETURNING *;


-- name: GetOrder :one
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
              WHERE c.id = p.category_id 
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
          WHERE cv.id = o.color_varian_id 
        )
    ) AS product,
    sv.size AS size,
    o.subtotal AS subtotal,
    o.quantity AS quantity,
    o.status AS status,
    o.created_at AS created_at,
    o.updated_at AS updated_at
 FROM orders o LEFT JOIN 
    products p ON o.product_id = p.id
    LEFT JOIN size_varians sv ON o.size_varian_id = sv.id
WHERE o.id = $1 LIMIT 1;

-- name: GetOrderForUpdate :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListOrder :many
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
              WHERE c.id = p.category_id 
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
          WHERE cv.id = o.color_varian_id 
        )
    ) AS product,
    sv.size AS size,
    o.subtotal AS subtotal,
    o.quantity AS quantity,
    o.status AS status,
    o.created_at AS created_at,
    o.updated_at AS updated_at
 FROM orders o LEFT JOIN 
    products p ON o.product_id = p.id
    LEFT JOIN size_varians sv ON o.size_varian_id = sv.id
ORDER BY o.id
LIMIT $1
OFFSET $2;

-- name: UpdateOrder :one
UPDATE orders
 set status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;