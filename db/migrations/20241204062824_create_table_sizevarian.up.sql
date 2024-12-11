CREATE TABLE "size_varians" (
  "id" bigserial PRIMARY KEY,
  "color_varian_id" bigint NOT NULL,
  "size" varchar NOT NULL,
  "stock" bigint NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz DEFAULT NULL
);


CREATE INDEX ON "size_varians" ("color_varian_id");

ALTER TABLE "size_varians" ADD FOREIGN KEY ("color_varian_id") REFERENCES "color_varians" ("id");