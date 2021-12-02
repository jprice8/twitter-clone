BEGIN;
CREATE TABLE IF NOT EXISTS "products" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "description" varchar(255) NOT NULL,
    "price" real NOT NULL,
    "created_at" timestamp default current_timestamp
);
COMMIT;