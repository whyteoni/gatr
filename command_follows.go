package main

import (
	"context"
	"fmt"

	"github.com/whyteoni/gatr/internal/database"
)

func CommandFollowing(state state, args []string, user database.User) (err error) {
	if len(args) != 0 {
		return fmt.Errorf("no args expected")
	}

	follows, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return
	}

	fmt.Printf("Following %d feeds.\n", len(follows))
	for _, f := range follows {
		fmt.Printf("  - %s [%s]\n", f.FeedName, f.FeedUrl)
	}
	return
}
