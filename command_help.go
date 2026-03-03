package main

import (
	"fmt"
	"math/rand"
	"strings"
)


var TagLines = []string{
	"Gatr's feeds better be using IP tables",
	"Gatr need his blogs, you SCSI-busted script kiddie",
	"Desk Pop! Now you can't hear. So read what Gatr cookin'",
}


func CommandHelp(state state, args []string) (err error) {
	const padding string = " "
	var offset int = 0

	tagLine := TagLines[rand.Intn(len(TagLines))]
	fmt.Printf("\n%s\n\n",tagLine)

	for _, command := range Commands {
		if len(command.Name) > offset {
			offset = len(command.Name)
		}
	}

	for _, command := range Commands {
		filler := strings.Repeat(padding, offset - len(command.Name))
		fmt.Printf("%s%s: %s\n", filler, command.Name, command.Desc)
	}
	fmt.Println()
	return
}
