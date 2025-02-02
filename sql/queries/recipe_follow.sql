-- name: GetUsersRecipes :many
SELECT * FROM recipe_follows
WHERE user_id = $1;

-- name: UnfollowRecipe :one
DELETE FROM recipe_follows
WHERE user_id = $1 AND recipe_id = $2
RETURNING *;

-- name: FollowRecipe :one
INSERT INTO recipe_follows (
    id, user_id, recipe_id
) VALUES ($1, $2, $3)
RETURNING *;
