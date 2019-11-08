package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	wio "github.com/Maki-Daisuke/go-whole-in-one"
)

func isValidName(name string) bool {
	re := regexp.MustCompile(`[^-(),.0-9=@A-Z[\]^_a-z{}~]`)
	return len(name) > 0 && !re.MatchString(name)
}

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		name = filepath.Base(cwd)
	}
	if !isValidName(name) {
		fmt.Fprintf(os.Stderr, "Invalid command name: %s\n", name)
		fmt.Fprintf(os.Stderr, "You can use only the following characters in command name: (),-.0-9=@A-Z[]^_a-z{}~\n")
		os.Exit(1)
	}

	err := wio.WriteMainGo(name, "0.0.1")
	if err != nil {
		if os.IsExist(err) {
			fmt.Fprintln(os.Stderr, `Initialization failed: main.go already exists. Run wio-init in an empty directory.`)
			os.Exit(1)
		} else {
			panic(err)
		}
	}

	err = wio.WritePackingList(name)
	if err != nil {
		if os.IsExist(err) {
			fmt.Fprintln(os.Stderr, `Initialization failed: packing-list already exists. Run wio-init in an empty directory.`)
			os.Exit(1)
		} else {
			panic(err)
		}
	}

	err = wio.WritePackGo([]byte{})
	if err != nil {
		panic(err)
	}
}
