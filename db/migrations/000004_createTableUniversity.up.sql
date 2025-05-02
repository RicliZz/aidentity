CREATE TABLE IF NOT EXISTS "university" (
    "ID" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(100) NOT NULL,
    CONSTRAINT unique_name UNIQUE ("name")
)