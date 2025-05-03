CREATE TABLE IF NOT EXISTS "profession" (
    "ID" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(50) UNIQUE NOT NULL
)