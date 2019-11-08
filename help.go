package wio

import (
	"fmt"
)

// HelpCommand is a predefined built-in subcommand, which shows
// the default help message.
var HelpCommand = FuncCommand(func(_ string, _ []string) {
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
