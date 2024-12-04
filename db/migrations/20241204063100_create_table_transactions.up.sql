
CREATE TABLE "transactions" (
  "tx_id" varchar PRIMARY KEY,
  "address_id" bigint NOT NULL,
  "shipping_id" bigint NOT NULL,
  "shipping_price" decimal NOT NULL,
  "total_price" decimal NOT NULL,
  "status" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "transactions" ("address_id");

CREATE INDEX ON "transactions" ("shipping_id");

CREATE INDEX ON "transactions" ("address_id", "shipping_id");

COMMENT ON COLUMN "transactions"."shipping_price" IS 'must be positive';

COMMENT ON COLUMN "transactions"."total_price" IS 'must be positive';

ALTER TABLE "transactions" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("shipping_id") REFERENCES "shippings" ("id");