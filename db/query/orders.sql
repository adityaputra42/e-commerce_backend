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
              WHERE c.id = p.category_id AND c.deleted_at IS NULL
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
          WHERE cv.id = o.color_varian_id AND cv.deleted_at IS NULL 
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
     products p ON o.product_id = p.id AND p.deleted_at IS NULL
     LEFT JOIN size_varians sv ON o.size_varian_id = sv.id AND sv.deleted_at IS NULL
     LEFT JOIN transactions tx ON o.transaction_id = tx.tx_id AND tx.deleted_at IS NULL
     WHERE o.id = $1 LIMIT 1 AND o.deleted_at IS NULL;

-- name: GetOrderForUpdate :one
SELECT * FROM orders
WHERE deleted_at IS NULL AND id = $1 LIMIT 1
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
              WHERE c.id = p.category_id AND c.deleted_at IS NULL
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
          WHERE cv.id = o.color_varian_id AND cv.deleted_at IS NULL
        )
    ) AS product,
    sv.size AS size,
    o.subtotal AS subtotal,
    o.quantity AS quantity,
    o.status AS status,
    o.created_at AS created_at,
    o.updated_at AS updated_at
 FROM orders o 
 LEFT JOIN products p ON o.product_id = p.id AND p.deleted_at IS NULL
 LEFT JOIN size_varians sv ON o.size_varian_id = sv.id AND sv.deleted_at IS NULL
 WHERE o.deleted_at IS NULL AND o.status = $1
ORDER BY o.id
LIMIT $2
OFFSET $3;

-- name: UpdateOrder :one
UPDATE orders
 set status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
UPDATE orders
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;