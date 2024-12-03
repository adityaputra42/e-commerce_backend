CREATE TABLE colorvarians (
  id bigint NOT NULL AUTO_INCREMENT,
  product_id bigint NOT NULL,
  color_code varchar(255) NOT NULL,
  images varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(product_id) REFERENCES products(id)
  );