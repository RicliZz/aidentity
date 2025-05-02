CREATE TABLE IF NOT EXISTS "student" (
    ID uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "userID" uuid REFERENCES "user" ("ID") ON DELETE CASCADE,
    "specialityID" uuid REFERENCES "speciality" ("ID") ON DELETE SET NULL
)