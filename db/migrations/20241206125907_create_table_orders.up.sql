CREATE TABLE "orders" (
  "id" varchar PRIMARY KEY,
  "transaction_id" varchar NOT NULL,
  "product_id" bigint NOT NULL,
  "color_varian_id" bigint NOT NULL,
  "size_varian_id" bigint NOT NULL,
  "unit_price" float NOT NULL,
  "subtotal" float NOT NULL,
  "quantity" bigint NOT NULL,
  "status" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz DEFAULT NULL
);
CREATE INDEX ON "orders" ("transaction_id");

CREATE INDEX ON "orders" ("product_id");

CREATE INDEX ON "orders" ("color_varian_id");

CREATE INDEX ON "orders" ("size_varian_id");

CREATE INDEX ON "orders" ("transaction_id", "product_id", "color_varian_id", "size_varian_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("tx_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("color_varian_id") REFERENCES "color_varians" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("size_varian_id") REFERENCES "size_varians" ("id");
