CREATE TABLE categories (
  id bigint NOT NULL AUTO_INCREMENT,
  product_id bigint NOT NULL,
  name varchar(255) NOT NULL,
  icon varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(uid) REFERENCES users(id)
  ) engine = InnoDB;