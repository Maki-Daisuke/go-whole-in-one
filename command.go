package wio

type Command interface {
	Exec(subname string, args []string)
}

type FuncCommand func(subname string, args []string)

func (f FuncCommand) Exec(subname string, args []string) {
	f(subname, args)
}
