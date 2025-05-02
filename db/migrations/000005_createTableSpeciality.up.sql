CREATE TABLE IF NOT EXISTS "speciality" (
    "ID" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "universityID" uuid REFERENCES "university" ("ID") ON DELETE CASCADE,
    name varchar(100) NOT NULL,
    code varchar(30) NOT NULL
)