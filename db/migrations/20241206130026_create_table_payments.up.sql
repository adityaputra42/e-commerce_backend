
CREATE TABLE "payment" (
  "id" bigserial PRIMARY KEY,
  "transaction_id" varchar NOT NULL,
  "total_payment" decimal NOT NULL,
  "status" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "payment" ("transaction_id");

CREATE INDEX ON "payment" ("transaction_id");

ALTER TABLE "payment" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("tx_id");
