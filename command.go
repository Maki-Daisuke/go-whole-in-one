package wio

// The Command type represents a built-in subcommand, which can be
// registered with Register function.
type Command interface {
	Exec(subname string, args []string)
}

// The FuncCommand type is an adapter to allow the use of ordinary functions
// as built-in subcommand.
// If f is a function with the appropriate signature, FuncCommand(f) is
// a Command that calls f.
type FuncCommand func(subname string, args []string)

// Exec calls f(subname, args).
func (f FuncCommand) Exec(subname string, args []string) {
	f(subname, args)
}
