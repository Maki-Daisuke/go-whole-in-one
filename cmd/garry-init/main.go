package main

import (
	"os"
	"path/filepath"

	"github.com/Maki-Daisuke/garry"
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
	err = garry.WriteMainGo(name, "0.0.1")
	if err != nil {
		panic(err)
	}
	err = garry.WritePackGo([]byte{})
	if err != nil {
		panic(err)
	}
	err = garry.WritePackingList(name)
	if err != nil {
		panic(err)
	}
}
