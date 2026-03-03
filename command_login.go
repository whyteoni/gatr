package main

import (
	"context"
	"fmt"
	"strings"
)

func CommandLogin(state state, args []string) (err error) {
	if len(args) != 1 {
		return fmt.Errorf("error: Gatr login one user at a time")
	}

	userName := args[0]
	_, err = state.db.GetUser(context.Background(), userName)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("%s is not registered", userName)
		}
		return
	}

	state.cfg.SetUser(userName)
	fmt.Printf("Gatr got you, %s.\n", userName)
	return
}
