package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/whyteoni/gatr/internal/config"
	"github.com/whyteoni/gatr/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	var args []string
	var cfg config.Config
	var err error
	var state state

	// Setup config
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error when reading config file: %v", err)
		os.Exit(1)
	}
	state.cfg = &cfg

	// Setup DB connections.
	db, err := sql.Open("postgres", state.cfg.DB_url)
	if err != nil {
		fmt.Printf("error when accessing database: %s", err)
		os.Exit(1)
	}
	state.db = database.New(db)

	// Validate command name and assign args
	if len(os.Args) < 2 {
		fmt.Println("error: No command given")
		CommandHelp(state, args)
		os.Exit(1)
	}
	commandName := os.Args[1]
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	// Get the function
	command, ok := Commands[commandName]
	if !ok {
		fmt.Printf("error: %s is not a valid command", commandName)
		CommandHelp(state, args)
		os.Exit(1)
	}

	// Run the function
	err = command.Callback(state, args)
	if err != nil {
		fmt.Printf("Gatr encountered: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
