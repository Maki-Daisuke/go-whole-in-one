package frontal

import "fmt"

var helpCmd = FuncCommand(func(_ string, _ []string) {
	fmt.Printf("usage: %s [--version] [--help] <subcommand> <args>\n", Name)
})
