package main

import (
	"os"
	"path/filepath"

	"github.com/Maki-Daisuke/uum7"
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
	err = uum7.WriteMainGo(name, "0.0.1")
	if err != nil {
		panic(err)
	}
	err = uum7.WritePackGo([]byte{})
	if err != nil {
		panic(err)
	}
	err = uum7.WritePackingList(name)
	if err != nil {
		panic(err)
	}
}
