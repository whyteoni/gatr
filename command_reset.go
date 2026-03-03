package main

import (
	"context"
	"fmt"
)

func CommandReset(state state, args []string) (err error) {
	if len(args) != 0 {
		return fmt.Errorf("no args expected")
	}

	err = state.db.ResetUsers(context.Background())
	if err == nil {
		fmt.Println("Gatr drained the swamp, all tadpoles out!")
	}
	return
}
