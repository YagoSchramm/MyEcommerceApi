SELECT id, product_id, user_id, value, quantity, created_at
FROM purchases
WHERE user_id = $1
ORDER BY created_at DESC;