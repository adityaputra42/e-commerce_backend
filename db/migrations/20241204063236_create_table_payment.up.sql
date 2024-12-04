
CREATE TABLE "payment" (
  "id" bigserial PRIMARY KEY,
  "payment_method_id" bigint NOT NULL,
  "transaction_id" varchar NOT NULL,
  "total_payment" decimal NOT NULL,
  "status" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "payment" ("payment_method_id");

CREATE INDEX ON "payment" ("transaction_id");

CREATE INDEX ON "payment" ("payment_method_id", "transaction_id");

COMMENT ON COLUMN "payment"."total_payment" IS 'must be positive';

ALTER TABLE "payment" ADD FOREIGN KEY ("payment_method_id") REFERENCES "payment_method" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("tx_id");