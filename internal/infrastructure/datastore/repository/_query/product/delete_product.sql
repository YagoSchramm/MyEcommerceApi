UPDATE products
SET deleted_at = NOW()
WHERE id=$1 AND user_id=$2;