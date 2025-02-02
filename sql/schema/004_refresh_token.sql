-- +goose Up
CREATE TABLE refresh_tokens(
    token VARCHAR(255) NOT NULL UNIQUE,
    valid_from TIMESTAMP NOT NULL,
    valid_until TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP,
    user_id INT NOT NULL REFERENCES users(id),
    PRIMARY KEY(token)
);

-- +goose Down
DROP TABLE refresh_tokens;
