-- +goose Up
CREATE TABLE feed_follows (
    id  uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    user_id     uuid NOT NULL
                REFERENCES users(id) 
                ON DELETE CASCADE,
    feed_id     uuid NOT NULL
                REFERENCES feeds(id) 
                ON DELETE CASCADE,
    CONSTRAINT user_and_feed_uniq_combo
    UNIQUE (user_id, feed_id) 
);

-- +goose Down
DELETE FROM feed_follows;
DROP TABLE feed_follows CASCADE;
