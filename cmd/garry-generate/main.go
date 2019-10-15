package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Maki-Daisuke/garry"
	"github.com/mattn/go-forlines"
)

func mustNotError(err error, msg string) {
	if err != nil {
		if msg == "" {
			msg = fmt.Sprintf("Error: %s", err)
		}
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}

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

var paths = strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))

func parsePackingList() []string {
	file, err := os.Open("./packing-list")
	mustNotError(err, "Can't find ./packing-list. Run `garry init` at first.")

	found := []string{}
	err = forlines.Do(file, func(line string) error {
		line = trimSpaceComment(line)
		if line == "" {
			return nil
		}
		if strings.Contains(line, "/") {
			found = append(found, findFile(line)...)
		} else {
			found = append(found, findCmd(line)...)
		}
		return nil
	})
	mustNotError(err, "")

	return found
}

var reTrim = regexp.MustCompile(`^\s*|\s*(?:#.*)$`)

func trimSpaceComment(s string) string {
	return reTrim.ReplaceAllLiteralString(s, "")
}

func findCmd(pattern string) []string {
	found := []string{}
	for _, path := range paths {
		m, err := filepath.Glob(filepath.Join(path, pattern))
		mustNotError(err, "")
		found = append(found, m...)
	}
	return found
}

func findFile(pattern string) []string {
	m, err := filepath.Glob(pattern)
	mustNotError(err, "")
	return m
}
