-- +goose Up
CREATE TABLE users (
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    name        text UNIQUE NOT NULL
);


-- +goose Down
DELETE FROM users;
DROP TABLE users CASCADE;
