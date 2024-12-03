CREATE TABLE payments (
  id bigint NOT NULL AUTO_INCREMENT,
  payment_method_id bigint NOT NULL,
  transaction_id varchar(255) NOT NULL,
  stock bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(payment_method_id) REFERENCES payment_methods(id)
  FOREIGN kEY(transaction_id) REFERENCES transactions(id)
  );