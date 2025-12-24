

-- purchasing_details
DROP INDEX IF EXISTS idx_pd_item_id;
DROP INDEX IF EXISTS idx_pd_purchasing_id;

-- purchasing
DROP INDEX IF EXISTS idx_purchasing_user_id;
DROP INDEX IF EXISTS idx_purchasing_supplier_id;

-- DROP TABLE (child → parent)

DROP TABLE IF EXISTS purchasing_details;
DROP TABLE IF EXISTS purchasing;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS suppliers;
DROP TABLE IF EXISTS users;
