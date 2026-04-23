SELECT id, user_id, user_name, purchase_id, product_id, rating, description, created_at, updated_at
FROM ratings
WHERE user_id=$1 AND deleted_at IS NULL
ORDER BY created_at DESC;