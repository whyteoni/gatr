package main

import (
	"context"

	"github.com/whyteoni/gatr/internal/database"
)

var Commands = make(map[string]CliCommand)

func registerCommand(name, desc string, command func(state, []string) error) {
	Commands[name] = CliCommand{
		Name:     name,
		Desc:     desc,
		Callback: command,
	}
}

func LoginWrapper(cmd func(s state, a []string, u database.User) error) func(s state, a []string) error {
	return func(s state, a []string) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser.Name)
		if err != nil {
			return err
		}
		return cmd(s, a, user)
	}
}

func init() {
	registerCommand("login", "USERNAME; Get in the swamp. Gatr gunna find you.", CommandLogin)
	registerCommand("help", "Pop your head up and take a look around.", CommandHelp)
	registerCommand("register", "USERNAME; Make your mark and pop that desk.", CommandRegister)
	registerCommand("users", "Listen up tadpoles, Roll Call!", CommandUsers)
	registerCommand("reset", "Gatr clear the swamp, all tadpoles out.", CommandReset)
	registerCommand("agg", "DURATION; Agg'Gregate them feeds. (Runs forever)", CommandAgg)
	registerCommand("addfeed", "NAME URL; Set Gatr tracking a new feed.", LoginWrapper(CommandAddFeed))
	registerCommand("feeds", "Checking cupboards.", CommandFeeds)
	registerCommand("follow", "URL; Follow a feed.", LoginWrapper(CommandFollow))
	registerCommand("following", "Check what you're following.", LoginWrapper(CommandFollowing))
	registerCommand("unfollow", "URL [URL...]; Stop following feeds", LoginWrapper(CommandUnfollow))
	registerCommand("browse", "[LIMIT]; Look at recent posts, default limit of 2", LoginWrapper(CommandBrowse))
}
