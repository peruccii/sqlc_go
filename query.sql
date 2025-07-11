-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: CreateProduct :exec
INSERT INTO products (name, price)
VALUES ($1, $2);


-- name: GetProduct :one
SELECT * FROM products 
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET name = $2, price = $3
WHERE id = $1;
