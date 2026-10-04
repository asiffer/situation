CREATE TABLE IF NOT EXISTS "agents" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "version" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT current_timestamp,
    "updated_at" TIMESTAMP NOT NULL DEFAULT current_timestamp
);