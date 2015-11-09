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


Functions
---------

### `func Exec()`

`Exec` searches a command implementation corresponding to the name of subcommand
and executes it. Because it may call `exec` systemcall, lines following Exec()
will never be executed:

```go
func main(){
  frontal.Exec()      // This may call `exec` systemcall inside,
  fmt.Println("???")  // thus, execution never reaches here.
}
```

This rule is also the case for built-in commands. `Exec` calls `os.Exit(0)`
when a built-in command is successfully returned. If you want to report error
from your built-in command, you should call `os.Exit` manually with non-zero integer.

### `func Register(subname string, cmd frontal.Command)`

`Register` binds `cmd` to subcommand name `subname`.

```go
frontal.Register("foo", frontal.FuncCommand(func(argv0 string, argv []string){
  fmt.Printf("You called '%s' subcommand with: %v", argv0, argv)
}))
```

Variables
---------

### `var Name string`

`Name` holds the name of root command. Initial value is the command name
extracted from `os.Args[0]`. This value is used by built-in `help` and `version`
subcommand. Also, it is consulted to look up executable files for subcommands.

You can overwrite the value to customize the behavior of Frontal. For example:

```
func main(){
  frontal.Name = "mysupercmd"
  frontal.Exec()  // searches "mysupercmd-*" even if the actual command's name
                  // is "supercmd".
}
```

### `var Version string`

`Version` holds version string of your command. Its default value is `"0"`.
This is used by built-in `version` subcommand.

Expected use case is using `-ldflags` command line option when you build your app:

```
$ go build -ldflags "-X github.com/Maki-Daisuke/frontal.Version=0.1"
$ ./your-cmd version
your-cmd version 0.1
```

This will replace the value of `Version` with `"0.1"` at build time. Yes, it is
too long to type everytime! Maybe, the following usage is more practical.
In your code:

```go
package main

import "github.com/Maki-Daisuke/frontal"

var version string

func main(){
  frontal.Version = version
  frontal.Exec()
}
```

Then, in the command line:

```
$ go build -ldflags "-X main.version=0.1"
$ ./your-cmd version
your-cmd version 0.1
```

It looks better.


License
-------

The Simplified BSD License (2-clause). See LICENSE file also.


Author
------

Daisuke (yet another) Maki
