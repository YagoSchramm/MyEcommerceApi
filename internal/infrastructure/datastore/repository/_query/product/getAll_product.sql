SELECT id, user_id, user_name, name, value, image, stock, avg_rating, total_reviews,description, created_at, updated_at
FROM products
WHERE deleted_at IS NULL
ORDER BY created_at DESC;