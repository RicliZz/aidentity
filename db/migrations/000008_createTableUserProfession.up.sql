CREATE TYPE preference_type AS ENUM ('liked', 'disliked', 'parent', 'dream');

CREATE TABLE IF NOT EXISTS "user_profession" (
    "ID" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "userID" uuid NOT NULL REFERENCES "user" ("ID"),
    "professionID" uuid NOT NULL REFERENCES "profession" ("ID"),
    "preference" preference_type NOT NULL
)