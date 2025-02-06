-- +goose Up
CREATE TABLE sc_keys(
    hash_key BYTEA NOT NULL,
    block_key BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL,
    valid_until TIMESTAMP NOT NULL,
    signing_revoked_at TIMESTAMP,
    PRIMARY KEY(hash_key, block_key)
);

-- +goose Down
DROP TABLE keys;
