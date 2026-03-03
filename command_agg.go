package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (rssFeed *RSSFeed, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil { return }

	req.Header.Set("User-Agent", "gatr")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { return }

	var rss RSSFeed
	rssFeed = &rss
	err = xml.Unmarshal(body, rssFeed)
	if err != nil { return }

	// Clean up titles and descriptions
	rss.Channel.Description = html.EscapeString(rss.Channel.Description)
	rss.Channel.Title = html.EscapeString(rss.Channel.Title)
	fmt.Printf("Fixed: %s\n", rss.Channel.Title)
	for _, item := range rss.Channel.Item {
		item.Title = html.EscapeString(item.Title)
		item.Description = html.EscapeString(item.Description)
		fmt.Printf("  fixed: %s\n", item.Title)
	}

	return
}

func CommandAgg(state state, args []string) (err error) {
	feedURL := "https://www.wagslane.dev/index.xml"                                         
	rss, err := fetchFeed(context.Background(), feedURL)
	if err != nil { return }

	text, err := json.MarshalIndent(rss, "", "  ")
	if err != nil { return }

	fmt.Printf("%s\n", text)
	return nil
}
