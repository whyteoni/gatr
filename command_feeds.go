package main

import (
	"context"
	"fmt"
	"whyteoni/gatr/internal/database"

	"github.com/google/uuid"
)

func CommandFeeds(state state, args []string) (err error){
	if len(args) != 0 {
		return fmt.Errorf("no args expected")
	}

	var feeds []database.Feed
	feeds, err = state.db.ListFeeds(context.Background())
	if err != nil { return }

	userMap := make(map[uuid.UUID]string)
	var username string
	for _, feed := range feeds {
		username, err = getUserNameByID(feed.UserID, &userMap, state)
		if err != nil {	return }
		fmt.Printf("- [%s](%s) from %s\n", feed.Name, feed.Url, username)
	}
	return
}

func getUserNameByID(user_id uuid.UUID, userMap *map[uuid.UUID]string, state state) (userName string, err error){
	// Check if user_id has already been seen
	userName, exists := (*userMap)[user_id]
	if exists { return }

	// Get username if needed, update userMap
	user, err := state.db.GetUserByID(context.Background(), user_id)
	userName = user.Name
	(*userMap)[user_id] = userName

	return
}



