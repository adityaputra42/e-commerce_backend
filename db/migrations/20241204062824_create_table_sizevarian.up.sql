CREATE TABLE "size_varians" (
  "id" bigserial PRIMARY KEY,
  "color_varian_id" bigint NOT NULL,
  "color" varchar NOT NULL,
  "images" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "size_varians" ("color_varian_id");

ALTER TABLE "size_varians" ADD FOREIGN KEY ("color_varian_id") REFERENCES "color_varians" ("id");
