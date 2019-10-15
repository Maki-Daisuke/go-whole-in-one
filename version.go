package uum7

import "fmt"

var versionCmd = FuncCommand(func(_ string, _ []string) {
	fmt.Printf("%s version %s\n", Name, Version)
})
