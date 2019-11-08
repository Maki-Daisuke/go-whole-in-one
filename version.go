package wio

import "fmt"

// VersionCommand is a predefined built-in subcommand, which
// shows the name and the version number of the command.
var VersionCommand = FuncCommand(func(_ string, _ []string) {
	fmt.Printf("%s version %s\n", Name, Version)
})
