CREATE TABLE orders (
  id varchar(100) NOT NULL UNIQUE,
  transaction_id varchar(100) NOT NULL,
  product_id bigint NOT NULL,
  color_varian_id bigint NOT NULL,
  size_varian_id bigint NOT NULL,
  shipping_price bigint NOT NULL,
  unit_price decimal(10,2) NOT NULL,
  subtotal decimal(10,2) NOT NULL,
  quantity bigint NOT NULL,
  status varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(transaction_id) REFERENCES transactions(id)
  FOREIGN kEY(product_id) REFERENCES products(id)
  FOREIGN kEY(color_varian_id) REFERENCES colorvarians(id)
  FOREIGN kEY(size_varian_id) REFERENCES sizevarians(id)
  );