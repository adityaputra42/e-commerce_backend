CREATE TABLE users (
  uid varchar(255) NOT NULL UNIQUE,
  username varchar(100)  NOT NULL,
  password varchar(100) NOT NULL,
  full_name varchar(255) NOT NULL,
  email varchar(100) UNIQUE NOT NULL,
  role varchar(100) NOT NULL,
  updated_at timestamp NOT NULL DEFAULT current_timestamp  ON UPDATE current_timestamp,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY (uid)  
) engine = InnoDB;



