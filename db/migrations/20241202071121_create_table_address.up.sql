CREATE TABLE address (
  id bigint NOT NULL AUTO_INCREMENT,
  uid varchar(255) NOT NULL,
  recipient_name varchar(255) NOT NULL,
  recipient_phone_number varchar(255) NOT NULL,
  province varchar(255) NOT NULL,
  city varchar(255) NOT NULL,
  district varchar(255) NOT NULL,
  village varchar(255) NOT NULL,
  postal_code varchar(255) NOT NULL,
  full_address varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(uid) REFERENCES users(id)
  ) engine = InnoDB;