// +build darwin !android,linux

package wio

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

func lookupCmds(pattern string) []string {
	found := make([]string, 0)
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		matches, err := filepath.Glob(filepath.Join(path, pattern))
		if err != nil {
			panic(err) // malformed pattern
		}
		for _, m := range matches {
			if isExecutable(m) {
				found = append(found, m)
			}
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

	u, err := user.Current()
	var mask uint32 = 0001
	st, ok := (fi.Sys()).(*syscall.Stat_t)
	if !ok {
		panic(fmt.Errorf("Can't get syscall.Stat_t of %s", path))
	}
	if strconv.FormatUint(uint64(st.Uid), 10) == u.Uid {
		mask = 0100
	} else {
		gid := strconv.FormatUint(uint64(st.Gid), 10)
		groups, err := u.GroupIds()
		if err != nil {
			panic(err)
		}
		for _, g := range groups {
			if g == gid {
				mask = 0010
				break
			}
		}
	}
	return uint32(fi.Mode())&mask != 0
}
