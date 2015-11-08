package frontal

type Command interface {
	Exec(args []string)
}

type FuncCommand func(args []string)

func (f FuncCommand) Exec(args []string) {
	f(args)
}
