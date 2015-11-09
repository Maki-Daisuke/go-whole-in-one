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


License
-------

The Simplified BSD License (2-clause). See LICENSE file also.


Author
------

Daisuke (yet another) Maki
