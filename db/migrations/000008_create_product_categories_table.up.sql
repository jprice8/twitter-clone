BEGIN;
CREATE TABLE IF NOT EXISTS "product_categories" (
    "id" serial PRIMARY KEY, 
    "product_id" integer NOT NULL,
    "category_id" integer NOT NULL
);
CREATE INDEX "product_categories_product_id_index" ON "product_categories" ("product_id");
CREATE INDEX "product_categories_category_id_index" ON "product_categories" ("category_id");
ALTER TABLE "product_categories"
    ADD CONSTRAINT "product_categories_product_id_foreign" FOREIGN KEY ("product_id") REFERENCES "products" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "product_categories"
    ADD CONSTRAINT "product_categories_category_id_foreign" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
COMMIT;