-- +goose Up
CREATE TABLE keys(
    token BYTEA NOT NULL,
    type VARCHAR(10),
    created_at TIMESTAMP,
    valid_until TIMESTAMP,
    signing_revoked_at TIMESTAMP,
    PRIMARY KEY(token)
);

-- +goose Down
DROP TABLE keys;
