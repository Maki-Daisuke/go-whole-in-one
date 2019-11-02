package main

import (
	"bytes"

	wio "github.com/Maki-Daisuke/go-whole-in-one"
)

func main() {
	files := parsePackingList()
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	err := pack(buf, files)
	mustNotError(err, "")
	err = wio.WritePackGo(buf.Bytes())
	mustNotError(err, "")
}
