CREATE TABLE "shippings" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" decimal NOT NULL,
  "state" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


COMMENT ON COLUMN "shippings"."price" IS 'must be positive';