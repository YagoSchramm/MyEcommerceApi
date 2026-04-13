INSERT INTO ratings (
    id, user_id, user_name, purchase_id, product_id, rating, description, created_at, updated_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);