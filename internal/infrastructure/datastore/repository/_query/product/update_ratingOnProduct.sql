UPDATE products
SET 
    avg_rating = sub.avg,
    total_reviews = sub.count,
    updated_at = NOW()
FROM (
    SELECT 
        product_id,
        COALESCE(AVG(rating), 0) as avg,
        COUNT(*) as count
    FROM ratings
    WHERE product_id = $1 AND deleted_at IS NULL
    GROUP BY product_id
) as sub
WHERE products.id = sub.product_id;