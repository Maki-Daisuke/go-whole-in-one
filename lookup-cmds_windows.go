package wio

import (
	"os"
	"path/filepath"
	"strings"
)

var exts []string

func init() {
	// This code was copied from https://golang.org/src/os/exec/lp_windows.go
	x := os.Getenv(`PATHEXT`)
	if x != "" {
		for _, e := range strings.Split(strings.ToLower(x), `;`) {
			if e == "" {
				continue
			}
			if e[0] != '.' {
				e = "." + e
			}
			exts = append(exts, e)
		}
	} else {
		exts = []string{".com", ".exe", ".bat", ".cmd"}
	}
}

func lookupExecutables(pattern string) []string {
	var matched []string
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		m, err := filepath.Glob(filepath.Join(path, pattern))
		if err != nil {
			panic(err) // malformed pattern
		}
		matched = append(matched, m...)
		for _, ext := range exts {
			m, err := filepath.Glob(filepath.Join(path, pattern+ext))
			if err != nil {
				panic(err) // malformed pattern
			}
			matched = append(matched, m...)
		}
	}

	var dedup []string
	check := make(map[string]bool)
	for _, m := range matched {
		if _, exists := check[m]; !exists {
			check[m] = true
			dedup = append(dedup, m)
		}
	}

	var found []string
	for _, m := range dedup {
		if isExecutable(m) {
			found = append(found, m)
		}
	}

	return found
}

func isExecutable(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}
	ext := filepath.Ext(path)
	for _, e := range exts {
		if ext == e {
			return true
		}
	}
	return false
}
