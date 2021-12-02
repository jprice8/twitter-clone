BEGIN;
CREATE TABLE IF NOT EXISTS "order_details" (
    "id" serial PRIMARY KEY,
    "order_id" integer NOT NULL,
    "product_id" integer NOT NULL,
    "price" real NOT NULL,
    "quantity" integer NOT NULL
);
CREATE INDEX "order_details_order_id" ON "order_details" ("order_id");
CREATE INDEX "order_details_product_id" ON "order_details" ("product_id");
ALTER TABLE "order_details"
    ADD CONSTRAINT "order_details_order_id_foreign" FOREIGN KEY ("order_id") REFERENCES "orders" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "order_details"
    ADD CONSTRAINT "order_details_product_id_foreign" FOREIGN KEY ("product_id") REFERENCES "products" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
COMMIT;