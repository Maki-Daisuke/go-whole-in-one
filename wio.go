/*
Package wio provides API to customize your command-line application
written using WIO (Whole-In-One to Go).

You don't need to care most of them, because wio command cares
on befalf of you.

See https://github.com/Maki-Daisuke/go-whole-in-one for more information.
*/
package wio

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Name holds the name of root command.
// This value is used for vriety of purposes.
// For example, it is used to make messages output by built-in help and
// version subcommand, to look up executables for subcommands, and also to
// determine the name of cache directory.
//
// You can overwrite the value to customize the behavior of WIO. For example:
//
//   func init(){
//       wio.Name = "newname"
//   }
//
// Note that you must set Name in init() function, since wio has alreay created
// cache directory and done preparation stuffs before main() function is called.
var Name = ""

// Version holds version string of your command. Its default value is `"0"`.
// This is used in built-in `version` subcommand and also used to determine
// the name of cache directory.
//
// As well as Name variable, you must set Version in init() function,
// before main() is called.
var Version = "0"

var builtins = map[string]Command{}

// Register binds cmd to name as a built-in subcommand.
// For example, if you register as follows:
//
//   wio.Register("foo", wio.FuncCommand(func(argv0 string, argv []string){
//       fmt.Printf("You called '%s' subcommand with: %v", argv0, argv)
//   }))
//
// Then, you can invoke this function in your command line like this:
//
//   $ yourcmd foo bar baz
//   You called 'foo' subcommand with: [bar baz]
func Register(name string, cmd Command) {
	builtins[name] = cmd
}

// Exec searches a command implementation corresponding to the name of subcommand
// and executes it. Because it may call exec systemcall, lines following Exec
// will never be executed:
//
//   func main(){
//       wio.Exec(os.Args[1:])  // This may call exec systemcall inside,
//       fmt.Println("???")     // thus, execution never reaches here.
//   }
//
// This rule is also the case for built-in commands. Exec calls os.Exit(0)
// when a built-in command is successfully finished.
// If you want to return status code from your built-in command,
// you need to call os.Exit with non-zero integer manually.
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
