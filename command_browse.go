package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/whyteoni/gatr/internal/database"
)

func CommandBrowse(state state, args []string, user database.User) (err error) {
	var limit int = 2
	if len(args) == 1 {
		limit, err = strconv.Atoi(args[0])
		if err != nil {
			return err
		}
	} else if len(args) > 1 {
		return fmt.Errorf("I can only accept a singular number as a limit")
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := state.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return
	}

	if len(posts) == 0 {
		fmt.Printf("\nNo new posts\n\n")
	} else {
		fmt.Println()
	}

	for _, post := range posts {
		fmt.Printf("TITLE: %s\n", post.Title)
		fmt.Printf("DESC: %s\n", post.Description.String)
		fmt.Printf("LINK: %s\n", post.Url)
		fmt.Printf("PUBDATE: %s\n", post.PublishedAt.Time)
		fmt.Println()
	}
	return
}
