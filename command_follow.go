package main

import (
	"context"
	"fmt"

	"github.com/whyteoni/gatr/internal/database"
)

func CommandFollow(state state, args []string, user database.User) (err error) {
	if len(args) != 1 {
		return fmt.Errorf("error: Gatr starts following one feed at a time")
	}
	feedURL := args[0]

	feed, err := state.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return
	}

	var feedFollowParams database.CreateFeedFollowParams
	feedFollowParams.FeedID = feed.ID
	feedFollowParams.UserID = user.ID
	follow, err := state.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return
	}

	fmt.Printf("You're now tracking %s\n", follow.FeedName)
	return
}
