package frontal

import "fmt"

var versionCmd = FuncCommand(func(_ []string) {
	fmt.Printf("%s version %s\n", Name, Version)
})
