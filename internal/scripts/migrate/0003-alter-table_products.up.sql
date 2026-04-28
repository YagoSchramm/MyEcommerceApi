ALTER TABLE products
ALTER COLUMN user_name TYPE TEXT
USING user_name::TEXT;
