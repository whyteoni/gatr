package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/whyteoni/gatr/internal/database"
)

func CommandAddFeed(state state, args []string, user database.User) error {

	// ARG validation
	if len(args) != 2 {
		return fmt.Errorf("Gatr needs that what and where, that's it.")
	}
	name, url := args[0], args[1]
	if strings.ToLower(url[:4]) != "http" {
		return fmt.Errorf("URL must start with http")
	}

	// Create new feed
	var feedParams database.CreateFeedParams
	feedParams.Name = name
	feedParams.Url = url
	feedParams.UserID = user.ID

	feed, err := state.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	err = CommandFollow(state, []string{url}, user)
	if err != nil {
		return err
	}

	fmt.Printf("Gatr tracking '%s' at %s\n", feed.Name, feed.Url)
	return nil
}
