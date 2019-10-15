package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mattn/go-forlines"
)

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
