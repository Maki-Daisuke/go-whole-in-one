package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	wio "github.com/Maki-Daisuke/go-whole-in-one"
	"github.com/mattn/go-forlines"
)

var paths = filepath.SplitList(os.Getenv("PATH"))

func parsePackingList() []string {
	file, err := os.Open("./packing-list")
	mustNotError(err, "Can't find ./packing-list. Run `wio init` at first.")

	found := []string{}
	err = forlines.Do(file, func(line string) error {
		line = trimSpaceComment(line)
		if line == "" {
			return nil
		}
		if strings.Contains(line, "/") || strings.Contains(line, string(filepath.Separator)) {
			m, err := filepath.Glob(line)
			mustNotError(err, "")
			found = append(found, m...)
		} else {
			m, err := wio.LookupExecutables(line)
			mustNotError(err, "")
			found = append(found, m...)
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
