package main

import (
	"os"
	"path/filepath"

	wio "github.com/Maki-Daisuke/go-whole-in-one"
)

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

	var err error
	err = wio.WriteMainGo(name, "0.0.1")
	if err != nil {
		panic(err)
	}
	err = wio.WritePackGo([]byte{})
	if err != nil {
		panic(err)
	}
	err = wio.WritePackingList(name)
	if err != nil {
		panic(err)
	}
}
