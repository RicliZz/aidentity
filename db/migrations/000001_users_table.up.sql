CREATE TYPE "role" AS ENUM('student', 'admin');

CREATE TABLE IF NOT EXISTS "Users" (
    "ID" bigint GENERATED ALWAYS AS IDENTITY,
    "First_name" varchar(30) NOT NULL ,
    "Last_name" varchar(30) NOT NULL ,
    "Email" varchar(255) NOT NULL ,
    "Password" varchar (255) NOT NULL,
    "Created_at" timestamp DEFAULT now(),
    "Role" "role",
    CONSTRAINT unique_email UNIQUE ("Email"),
    CONSTRAINT valid_email CHECK ( "Email" LIKE '%@%' )
);