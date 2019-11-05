package wio

import (
	"fmt"
)

var helpCmd = FuncCommand(func(_ string, _ []string) {
	cmds, err := ListSubcommands()
	if err != nil {
		panic(err)
	}
	fmt.Printf("usage: %s SUBCOMMAND [ARGS...]\n\n", Name)
	fmt.Println("Available subcommands:")
	for _, c := range cmds {
		fmt.Printf("  %s\n", c)
	}
})
