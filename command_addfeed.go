package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"whyteoni/gatr/internal/database"

	"github.com/google/uuid"
)

func CommandAddFeed(state state, args []string) error {
	
	// ARG validation
	if len(args) != 2 {
		return fmt.Errorf("Gatr needs that what and where, that's it.")
	}
	name, url := args[0], args[1]
	if strings.ToLower(url[:4]) != "http" {
		return fmt.Errorf("URL must start with http")
	} 

	//Get User ID from DB
	currentUser := state.cfg.CurrentUserName
	if currentUser == "" {
		return fmt.Errorf("Gotta creep, creep a little less. Login first.")
	}
	user, err := state.db.GetUser(context.Background(), currentUser)
	if err != nil { return err }

	// Create new feed
	var feedParams database.CreateFeedParams
	feedParams.ID = uuid.New()
	feedParams.Name = name
	feedParams.Url = url
	feedParams.CreatedAt = time.Now()
	feedParams.UpdatedAt = time.Now()
	feedParams.UserID = user.ID

	feed, err := state.db.CreateFeed(context.Background(), feedParams)
	if err != nil { return err }

	fmt.Printf("Gatr tracking '%s' at %s\n", feed.Name, feed.Url)
	return nil
}
