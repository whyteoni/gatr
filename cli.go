package main

var Commands = make(map[string]CliCommand)
func registerCommand(name, desc string, command func(state, []string) error) {
	Commands[name] = CliCommand{
		Name: name,
		Desc: desc,
		Callback: command,
	}
}

func init() {
	registerCommand("login", "USERNAME; Get in the swamp. Gatr gunna find you.", CommandLogin)
	registerCommand("help","Pop your head up and take a look around.", CommandHelp)
	registerCommand("register", "USERNAME; Make your mark and pop that desk.", CommandRegister)
	registerCommand("users", "Listen up tadpoles, Roll Call!", CommandUsers)
	registerCommand("reset", "Gatr clear the swamp, all tadpoles out.", CommandReset)
	registerCommand("agg", "Gregate them feeds.", CommandAgg)
	registerCommand("addfeed", "NAME URL; Set Gatr tracking a new feed.", CommandAddFeed)
	registerCommand("feeds", "Checking cupboards.", CommandFeeds)
}
