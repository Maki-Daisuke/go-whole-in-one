package main

import (
	"bytes"

	"github.com/Maki-Daisuke/garry"
)

func main() {
	files := parsePackingList()
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	err := pack(buf, files)
	mustNotError(err, "")
	err = garry.WritePackGo(buf.Bytes())
	mustNotError(err, "")
}
