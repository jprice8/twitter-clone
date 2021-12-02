BEGIN;
CREATE TABLE IF NOT EXISTS "orders" (
    "id" serial PRIMARY KEY,
    "user_id" integer NOT NULL,
    "extended_cost" real NOT NULL,
    "created_at" timestamp default current_timestamp
);
CREATE INDEX "orders_user_id_index" ON "orders" ("user_id");
ALTER TABLE "orders"
    ADD CONSTRAINT "orders_user_id_foreign" FOREIGN KEY ("user_id") REFERENCES "users" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
COMMIT;