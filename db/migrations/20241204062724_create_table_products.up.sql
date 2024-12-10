
CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "category_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "images" varchar NOT NULL,
  "rating" float NOT NULL DEFAULT 0,
  "price" float NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "products" ("category_id");

COMMENT ON COLUMN "products"."price" IS 'must be positive';

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
