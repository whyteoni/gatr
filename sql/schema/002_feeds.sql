-- +goose Up
CREATE TABLE feeds (
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    name        text UNIQUE NOT NULL,
    url         text UNIQUE NOT NULL,
    user_id     uuid NOT NULL
                REFERENCES users(id) 
                ON DELETE CASCADE
);

-- +goose Down
DELETE FROM feeds;
DROP TABLE feeds CASCADE;
