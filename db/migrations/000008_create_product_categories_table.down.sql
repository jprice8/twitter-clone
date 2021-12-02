BEGIN;
ALTER TABLE IF EXISTS product_categories
    DROP CONSTRAINT IF EXISTS product_categories_product_id_foreign;
ALTER TABLE IF EXISTS product_categories
    DROP CONSTRAINT IF EXISTS product_categories_category_id_foreign;
DROP TABLE IF EXISTS product_categories;
COMMIT;