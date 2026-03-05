-- +goose Up
CREATE TABLE posts (
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at      timestamptz NOT NULL DEFAULT now(),
    updated_at      timestamptz NOT NULL DEFAULT now(),
    title           text NOT NULL,
    url             text UNIQUE NOT NULL,
    description     text,
    published_at    timestamptz,
    feed_id         uuid NOT NULL
                    REFERENCES feeds(id)
                    ON DELETE CASCADE
);

-- +goose Down
DELETE FROM posts;
DROP TABLE posts CASCADE;
