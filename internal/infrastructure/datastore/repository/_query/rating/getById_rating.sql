SELECT id, user_id, user_name, purchase_id, product_id, rating, description, created_at, updated_at, deleted_at
FROM ratings
WHERE id=$1 AND deleted_at IS NULL;