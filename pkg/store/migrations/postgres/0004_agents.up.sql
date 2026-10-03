CREATE TABLE IF NOT EXISTS "agents" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "version" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);