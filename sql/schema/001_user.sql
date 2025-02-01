-- +goose Up
CREATE TABLE users(
    id INT NOT NULL,
    username VARCHAR(64),
    password VARCHAR(64),
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users;
