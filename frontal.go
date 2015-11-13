package frontal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var Name = filepath.Base(os.Args[0])
var Version = "0"
var builtins = map[string]Command{}

func Register(name string, cmd Command) {
	builtins[name] = cmd
}

func init() {
	Register("help", helpCmd)
	Register("--help", helpCmd)
	Register("version", versionCmd)
	Register("--version", versionCmd)
}

func Exec() {
	if len(os.Args) <= 1 {
		builtins["--help"].Exec("--help", []string{})
		os.Exit(1)
	}
	subname := os.Args[1]
	if cmd, ok := builtins[subname]; ok {
		cmd.Exec(subname, os.Args[2:])
		os.Exit(0)
	}
	cmdname := Name + "-" + subname
	if _, err := exec.LookPath(cmdname); err != nil {
		fmt.Printf("%s: '%s' is not a %s command. See '%s --help'.\n", Name, subname, Name, Name)
		os.Exit(1)
	}
	os.Args[1] = cmdname
	syscall.Exec(cmdname, os.Args[1:], []string{})
	// If you are here, exec systemcall failed.
	// We'll fallback to os/exec module for non-exec-able platfaorm, e.g. Windows.
	cmd := exec.Command(cmdname, os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
