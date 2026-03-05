-- +goose Up
ALTER TABLE feeds
ADD last_fetched_at timestamptz;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_fetched_at CASCADE;
