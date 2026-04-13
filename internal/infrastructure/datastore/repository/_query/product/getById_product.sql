SELECT id, user_id, user_name, name, value, image, avg_rating, total_reviews, stock, description, created_at, updated_at, deleted_at
FROM products
WHERE id=$1 AND deleted_at IS NULL;