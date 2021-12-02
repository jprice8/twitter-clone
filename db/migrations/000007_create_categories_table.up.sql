BEGIN;
CREATE TABLE IF NOT EXISTS "categories" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "description" varchar(255) NOT NULL
);
COMMIT;