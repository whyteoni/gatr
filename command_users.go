package main

import (
	"context"
	"fmt"
	"whyteoni/gatr/internal/database"
)

func CommandUsers(state state, args []string) (err error) {
	if len(args) != 0 {
		return fmt.Errorf("no args expected")
	}

	var users []database.User
	users, err = state.db.ListUsers(context.Background())
	if err != nil { return }

	for _, user := range users {
		marker := ""
		if user.Name == state.cfg.CurrentUserName { marker = " (current)" }
		fmt.Printf("* %s%s\n", user.Name, marker )
	}
	return
}
