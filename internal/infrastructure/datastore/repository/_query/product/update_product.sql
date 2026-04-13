UPDATE products
SET name=$1, value=$2, image=$3, stock=$4, description=$5, updated_at=$6
WHERE id=$7 AND deleted_at IS NULL;