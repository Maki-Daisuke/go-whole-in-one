package main

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/Maki-Daisuke/garry"
)

func main() {
	files := parsePackingList()
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	args := append([]string{"zcf", "-"}, files...)
	cmd := exec.Command("tar", args...)
	cmd.Stdin = nil
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	mustNotError(err, "")

	err = garry.WritePackGo(buf.Bytes())
	mustNotError(err, "")
}
