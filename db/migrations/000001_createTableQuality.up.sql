CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "quality" (
    "ID" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "nameQuality" varchar(50) NOT NULL,
    CONSTRAINT "unique_nameQuality" UNIQUE ("nameQuality")
)