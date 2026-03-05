package main

import (
	"context"
	"fmt"

	"github.com/whyteoni/gatr/internal/database"
)

func CommandUnfollow(state state, args []string, user database.User) (err error) {
	var params database.UnfollowParams
	params.UserID = user.ID
	for _, feedURL := range args {
		feed, err := state.db.GetFeed(context.Background(), feedURL)
		if err != nil {
			return err
		}

		params.FeedID = feed.ID
		err = state.db.Unfollow(context.Background(), params)
		if err != nil {
			return err
		}
		fmt.Printf("Unfollowed: %s\n", feed.Name)
	}
	return nil
}
