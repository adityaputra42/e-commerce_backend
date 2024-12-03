CREATE TABLE payment_methods (
  id bigint NOT NULL AUTO_INCREMENT,
  account_name varchar(255) NOT NULL,
  account_number varchar(255) NOT NULL,
  bank_name varchar(255) NOT NULL,
  bank_icons varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  );