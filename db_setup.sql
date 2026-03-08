CREATE DATABASE gatr;

\c gatr;

CREATE TABLE users (
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    name        text UNIQUE NOT NULL
);

CREATE TABLE feeds (
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at      timestamptz NOT NULL DEFAULT now(),
    updated_at      timestamptz NOT NULL DEFAULT now(),
    name            text UNIQUE NOT NULL,
    url             text UNIQUE NOT NULL,
    user_id         uuid NOT NULL
                    REFERENCES users(id) 
                    ON DELETE CASCADE,
    last_fetched_at timestamptz
);

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
