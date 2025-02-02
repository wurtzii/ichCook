-- +goose Up
CREATE TABLE recipe_follows(
    id INT NOT NULL,
    recipe_id INT NOT NULL REFERENCES recipes(id),
    user_id INT NOT NULL REFERENCES users(id),
    PRIMARY KEY(recipe_id, user_id)
);

-- +goose Down
DROP TABLE recipe_follows;
