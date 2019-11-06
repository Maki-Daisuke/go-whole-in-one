package wio

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

var Name = ""
var Version = "0"
var builtins = map[string]Command{}

func Register(name string, cmd Command) {
	builtins[name] = cmd
}

func init() {
	Register("help", helpCmd)
	Register("-h", helpCmd)
	Register("--help", helpCmd)
	Register("version", versionCmd)
	Register("-v", versionCmd)
	Register("--version", versionCmd)
}

func Exec(args []string) {
	if len(args) == 0 {
		builtins["--help"].Exec("--help", []string{})
		os.Exit(1)
	}
	subname := args[0]
	if cmd, ok := builtins[subname]; ok {
		cmd.Exec(subname, args[1:])
		os.Exit(0)
	}
	cmdname := Name + "-" + subname
	if _, err := exec.LookPath(cmdname); err != nil {
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", Name, subname, Name, Name)
		os.Exit(1)
	}
	args[0] = cmdname
	syscall.Exec(cmdname, args, os.Environ())
	// If you are here, exec systemcall failed.
	// We'll fallback to os/exec module for non-exec-able platfaorm, e.g. Windows.
	cmd := exec.Command(cmdname, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
