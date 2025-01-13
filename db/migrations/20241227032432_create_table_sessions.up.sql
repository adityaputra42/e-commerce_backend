CREATE TABLE "user_sessions" (
  "id" UUID NOT NULL UNIQUE PRIMARY KEY,
  "user_uid" VARCHAR NOT NULL,
  "refresh_token" VARCHAR NOT NULL,
  "user_agent" VARCHAR NOT NULL,
  "client_ip" VARCHAR NOT NULL,
  "is_blocked" BOOLEAN NOT NULL DEFAULT false,
  "expired_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE "user_sessions" 
ADD FOREIGN KEY ("user_uid") REFERENCES "users" ("uid");