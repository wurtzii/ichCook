-- +goose Up
CREATE TABLE recipes(
    id INT NOT NULL,
    user_id INT NOT NULL REFERENCES users(id),
    ingredients TEXT,
    instructions TEXT,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE recipes;
