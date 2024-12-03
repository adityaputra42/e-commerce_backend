CREATE TABLE sizevarians (
  id bigint NOT NULL AUTO_INCREMENT,
  color_varian_id bigint NOT NULL,
  size varchar(10) NOT NULL,
  stock bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp,
  updated_at timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at timestamp NULL,
  PRIMARY KEY(id),
  FOREIGN kEY(color_varian_id) REFERENCES colorvarians(id)
  );