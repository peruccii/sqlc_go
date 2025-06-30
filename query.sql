-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: CreateProduct :exec
INSERT INTO products (name, price)
VALUES ($1, $2);
