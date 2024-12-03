CREATE TABLE transactions (
  id varchar(100) NOT NULL UNIQUE,
  address_id bigint NOT NULL,
  shipping_id bigint NOT NULL,
  shipping_price bigint NOT NULL,
  status varchar(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(address_id) REFERENCES address(id)
  FOREIGN kEY(shipping_id) REFERENCES shippings(id)
  );