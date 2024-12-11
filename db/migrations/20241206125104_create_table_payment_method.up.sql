
CREATE TABLE "payment_method" (
  "id" bigserial PRIMARY KEY,
  "account_name" varchar NOT NULL,
  "account_number" varchar NOT NULL,
  "bank_name" varchar NOT NULL,
  "bank_images" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz DEFAULT NULL
);

