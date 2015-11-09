Frontal
=======

Yet another Go library for Git way of subcommand.


Synopsis
--------

You need to just call `frontal.Exec`:

```go
package main

import "github.com/Maki-Daisuke/frontal"

func main(){
  frontal.Exec()
}
```

Then, you can call your command as follows:

```
$ yourcmd sub arg1 arg2
```

It searches `yourcmd-sub` with consulting `PATH` environment variable and
execute it passing rest of the arguments.


Description
-----------

This is yet another Go library for Git way of subcommand. This is inspired by
[façade](https://github.com/pepabo/facade) and I like its concept. However,
I don't need fancy colored output and it is even harmful when I implement filter
commands, which read from STDIN, do something and write output to STDOUT.

I wanted more simple and low I/O overhead library. The less overhead is the better,
of course. That's why I wrote Frontal.


Built-in Command
----------------

You can implement your subcommand directly in Go by implementing `frontal.Command`
interface:

```go
type Command interface {
	Exec(subname string, args []string)
}
```

For your convenience, you can use `FuncCommand` type. For example:

```go
func main(){
  frontal.Register("mycmd", frontal.FuncCommand(func(argv0 string, args []string){
    fmt.Println("This is my command!")
  }))
  frontal.Exec()
}
```

Frontal has the following predefined built-in subcommands:

- `help`
  - Shows simple help message
- `--help`
  - Alias of `help`
- `version`
  - Shows version number
- `--version`
  - Alias of `version`

You can easily override them calling `frontal.Register`.


License
-------

The Simplified BSD License (2-clause). See LICENSE file also.


Author
------

Daisuke (yet another) Maki
