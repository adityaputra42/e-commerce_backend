

CREATE TABLE "address" (
  "id" bigserial PRIMARY KEY,
  "uid" varchar NOT NULL,
  "recipient_name" varchar NOT NULL,
  "recipient_phone_number" varchar NOT NULL,
  "province" varchar NOT NULL,
  "city" varchar NOT NULL,
  "district" varchar NOT NULL,
  "village" varchar NOT NULL,
  "postal_code" varchar NOT NULL,
  "full_address" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "address" ("uid");

CREATE UNIQUE INDEX ON "address" ("uid");

ALTER TABLE "address" ADD FOREIGN KEY ("uid") REFERENCES "users" ("uid");