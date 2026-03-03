package main

import (
	"context"
	"fmt"
	"time"

	"whyteoni/gatr/internal/database"

	"github.com/google/uuid"
)

func CommandRegister(state state, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Gatr registers one user at a time.")
	}

	var userParams database.CreateUserParams
	userParams.Name = args[0]
	userParams.ID = uuid.New()
	userParams.CreatedAt = time.Now()
	userParams.UpdatedAt = time.Now()

	user, err := state.db.CreateUser(context.Background(), userParams)
	if err != nil {	return err	}

	fmt.Printf("Gatr registered %s (%s) at %s\n", user.Name, user.ID, user.CreatedAt)
	if err = CommandLogin(state, []string{user.Name}); err != nil { return err }
	return nil

}
