package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/whyteoni/gatr/internal/database"
)

var timeLayouts = []string{
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02 15:04:05",
	"2006-01-02",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02T15:04:05",
}

func fetchFeed(ctx context.Context, feedURL string) (rssFeed *RSSFeed, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "gatr")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return
	}

	// Clean up titles and descriptions
	rssFeed.Channel.Description = html.EscapeString(rssFeed.Channel.Description)
	rssFeed.Channel.Title = html.EscapeString(rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		item.Title = html.EscapeString(item.Title)
		item.Description = html.EscapeString(item.Description)
	}

	return
}

func fetchFeeds(state state) (err error) {
	feed, err := state.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return
	}

	feed, err = state.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return
	}

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return
	}

	for _, item := range rss.Channel.Item {
		params := database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: NewNullString(item.Description),
			PublishedAt: NewNullTime(item.PubDate),
			FeedID:      feed.ID,
		}

		post, err := state.db.CreatePost(context.Background(), params)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				continue
			}
			return err
		}
		fmt.Printf("Reaped '%s'(%s) for %s\n", post.Title, post.ID, feed.Name)
	}

	return
}

func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullTime(s string) sql.NullTime {
	if s == "" {
		return sql.NullTime{Valid: false}
	}
	t, err := ParseTime(s)
	if err != nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func ParseTime(value string) (time.Time, error) {
	var err error

	for _, layout := range timeLayouts {
		var t time.Time
		t, err = time.Parse(layout, value)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unsupported time format: %q", value)
}

func CommandAgg(state state, args []string) (err error) {
	if len(args) != 1 {
		return fmt.Errorf("Gatr keeps a schedule tight! Tell me how long between checkups.")
	}

	dwellTime, err := time.ParseDuration(args[0])
	if err != nil {
		return
	}
	fmt.Printf("Collections are every %s\n", dwellTime)

	ticker := time.NewTicker(dwellTime)
	for ; ; <-ticker.C {
		err = fetchFeeds(state)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}
	}
}
