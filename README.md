Frontal
=======

Yet another Go library for Git way of subcommand.


SYNOPSIS
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
$ yourcmd sub
```

It searches `yourcmd-sub` and runs it with consulting `PATH` env variable.
