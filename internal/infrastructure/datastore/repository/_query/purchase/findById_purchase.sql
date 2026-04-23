SELECT id, product_id, user_id, value, quantity, created_at
FROM purchases
WHERE id = $1;