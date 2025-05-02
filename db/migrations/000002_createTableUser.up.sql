CREATE TYPE "role" AS ENUM('schoolboy', 'student', 'adult');

CREATE TABLE IF NOT EXISTS "user" (
    "ID" uuid PRIMARY KEY default gen_random_uuid(),
    "fullName" varchar(255),
    "dateOfBirth" date,
    "role" role,
    "email" varchar(255) NOT NULL,
    "telegram" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    CONSTRAINT unique_email UNIQUE ("email"),
    CONSTRAINT unique_telegram UNIQUE ("telegram")
)