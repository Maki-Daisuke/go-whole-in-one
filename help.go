package frontal

import "fmt"

var helpCmd = FuncCommand(func(args []string) {
	fmt.Printf("usage: %s [--version] [--help] <subcommand> <args>\n", Name)
})
