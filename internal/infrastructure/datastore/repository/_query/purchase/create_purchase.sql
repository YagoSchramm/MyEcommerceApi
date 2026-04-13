INSERT INTO purchases (
    id, product_id, user_id, value, quantity, created_at
) VALUES ($1, $2, $3, $4, $5, $6);