-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (user_id, feed_id)
    VALUES ($1,$2)
    RETURNING *
)
SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users on inserted_feed_follow.user_id = users.id
INNER JOIN feeds on inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT 
    users.name AS user_name,
    feeds.name AS feed_name,
    feeds.url AS feed_url
FROM feed_follows
INNER JOIN users ON users.id = feed_follows.user_id
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE users.id = $1;

-- name: Unfollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;
