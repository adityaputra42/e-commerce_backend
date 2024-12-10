CREATE TABLE "transactions" (
  "tx_id" varchar PRIMARY KEY,
  "address_id" bigint NOT NULL,
  "shipping_id" bigint NOT NULL,
  "payment_method_id" bigint NOT NULL,
  "shipping_price" float NOT NULL,
  "total_price" float NOT NULL,
  "status" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "transactions" ("address_id");

CREATE INDEX ON "transactions" ("shipping_id");

CREATE INDEX ON "transactions" ("payment_method_id");

CREATE INDEX ON "transactions" ("address_id", "shipping_id", "payment_method_id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("shipping_id") REFERENCES "shippings" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("payment_method_id") REFERENCES "payment_method" ("id");
