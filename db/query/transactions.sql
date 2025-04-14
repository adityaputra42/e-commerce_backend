-- name: CreateTransaction :one
INSERT INTO transactions (
 tx_id,
 address_id,
 shipping_id,
 payment_method_id,
 shipping_price,
 total_price,
 status
) VALUES (
  $1, $2 ,$3, $4, $5, $6, $7
)
RETURNING *;


-- name: GetTransaction :one
SELECT 
      t.tx_id as tx_id,
            jsonb_build_object(
                  'id', a.id,
                  'recipient_name', a.recipient_name,
                  'recipient_phone_number', a.recipient_phone_number,
                  'province',a.province,
                  'city',a.city,
                  'district',a.district,
                  'village',a.village,
                  'postal_code',a.postal_code,
                  'full_address',a.full_address      
            ) AS address,
            jsonb_build_object(
                  'id',s.id,
                  'name',s.name,
                  'price',s.price,
                  'state',s.state,
                  'updated_at',s.updated_at,
                  'created_at',s.created_at
            ) AS shipping,
            jsonb_build_object(
                  'id',pm.id,
                  'account_name',pm.account_name,
                  'account_number',pm.account_number,
                  'bank_name',pm.bank_name,
                  'bank_images',pm.bank_images,
                  'updated_at',pm.updated_at,
                  'created_at',pm.created_at
            ) AS payment_method,
            (
            SELECT jsonb_agg(
                  jsonb_build_object(
                        'id',o.id,
                        'transaction_id', o.transaction_id,
                        'product',jsonb_build_object(
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
                        ), 
                        'size',sv.size,
                        'subtotal',o.subtotal,
                        'quantity',o.quantity,
                        'status',o.status,
                        'created_at',o.created_at,
                        'updated_at',o.updated_at
                  ) 
            ) FROM orders o
              LEFT JOIN products p ON o.product_id = p.id AND p.deleted_at IS NULL
              LEFT JOIN size_varians sv ON o.size_varian_id = sv.id AND sv.deleted_at IS NULL
              WHERE o.transaction_id = t.tx_id AND o.deleted_at IS NULL
            ) AS orders,
      t.shipping_price AS shipping_price,
      t.total_price AS total_price,
      t.status AS status,
      t.created_at AS created_at,
      t.updated_at AS updated_at
 FROM transactions t 
 LEFT JOIN address a ON t.address_id = a.id AND a.deleted_at IS NULL
 LEFT JOIN shippings s ON t.shipping_id = s.id AND s.deleted_at IS NULL
 LEFT JOIN payment_method pm ON t.payment_method_id = pm.id AND pm.deleted_at IS NULL
 WHERE t.deleted_at IS NULL AND t.tx_id = $1 LIMIT 1;

-- name: GetTransactionForUpdate :one
SELECT * FROM transactions
WHERE deleted_at IS NULL AND tx_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTransaction :many
SELECT 
      t.tx_id as tx_id,
      t.address_id as address_id,
      t.payment_method_id as payment_method_id,
      t.shipping_id as shipping_id,
      t.shipping_price AS shipping_price,
      t.total_price AS total_price,
      t.status AS status,
      t.created_at AS created_at,
      t.updated_at AS updated_at,
      COUNT(o.id) AS total_orders
FROM transactions t
LEFT JOIN orders o ON t.tx_id = o.transaction_id AND o.deleted_at IS NULL
WHERE t.deleted_at IS NULL
GROUP BY t.tx_id
ORDER BY t.tx_id
LIMIT $1
OFFSET $2;

-- name: UpdateTransaction :one
UPDATE transactions
 set status = $2
WHERE tx_id = $1
RETURNING *;

-- name: DeleteTransaction :exec
UPDATE transactions
SET deleted_at = CURRENT_TIMESTAMP
WHERE tx_id = $1;