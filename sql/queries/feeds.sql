-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES ($1,$2,$3)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1;

-- name: ListFeeds :many
SELECT * FROM feeds;

-- name: ResetFeeds :exec
TRUNCATE TABLE feeds;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = now(), updated_at = now()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;
