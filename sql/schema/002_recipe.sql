-- +goose Up
CREATE TABLE recipes(
    id INT NOT NULL,
    ingredients TEXT,
    instructions TEXT,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE recipes;
