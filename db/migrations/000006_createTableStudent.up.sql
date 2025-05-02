CREATE TABLE IF NOT EXISTS "student" (
    "ID" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "userID" uuid REFERENCES "user" ("ID") ON DELETE CASCADE,
    "studyYear" int NOT NULL CHECK ("studyYear" BETWEEN 1 AND 10),
    "specialityID" uuid REFERENCES "speciality" ("ID") ON DELETE SET NULL
)