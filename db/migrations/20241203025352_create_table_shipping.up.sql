CREATE TABLE shippings (
  id bigint NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  state varchar(50) NOT NULL,
  price decimal(10,2) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  );