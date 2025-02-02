-- name: GetRecipeByID :one
SELECT * FROM recipes
WHERE id = $1 LIMIT 1;

-- name: DeleteRecipeById :one
DELETE FROM recipes
WHERE id = $1
RETURNING *;
