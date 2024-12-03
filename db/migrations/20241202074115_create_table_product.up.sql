CREATE TABLE products (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name varchar(100)  NOT NULL,
  description varchar(255) NOT NULL,
  image varchar(255) NOT NULL,
  rating decimal(10,2)  NOT NULL,
  price decimal(10,2) NOT NULL,
  updated_at timestamp NOT NULL DEFAULT current_timestamp  ON UPDATE current_timestamp,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY (id)  
);


