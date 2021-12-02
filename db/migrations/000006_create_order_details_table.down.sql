BEGIN;
ALTER TABLE IF EXISTS order_details
    DROP CONSTRAINT IF EXISTS order_details_order_id_foreign;
ALTER TABLE IF EXISTS order_details
    DROP CONSTRAINT IF EXISTS order_details_product_id_foreign;
DROP TABLE IF EXISTS order_details;
COMMIT;