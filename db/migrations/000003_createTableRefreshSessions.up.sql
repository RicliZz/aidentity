CREATE TABLE If NOT EXISTS "refreshSession" (
    "ID" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "userID" uuid REFERENCES "user" ("ID") ON DELETE CASCADE,
    "refreshToken" uuid NOT NULL,
    "ua" varchar(200) NOT NULL,
    "fingerprint" varchar(200) NOT NULL,
    "IP" varchar(15) NOT NULL,
    "exp" bigint NOT NULL,
    "createdAt" TIMESTAMP NOT NULL DEFAULT now()
)