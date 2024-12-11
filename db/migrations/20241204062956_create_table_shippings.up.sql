CREATE TABLE "shippings" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" float NOT NULL,
  "state" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz DEFAULT NULL
);

COMMENT ON COLUMN "shippings"."price" IS 'must be positive';