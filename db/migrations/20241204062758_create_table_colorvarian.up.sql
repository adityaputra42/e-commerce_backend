CREATE TABLE "color_varians" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "color" varchar NOT NULL,
  "images" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "color_varians" ("product_id");

ALTER TABLE "color_varians" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");