package wio

import (
	"fmt"
	"os"
)

func setEnv(path string) {
	if lp := os.Getenv("DYLD_LIBRARY_PATH"); lp != "" {
		os.Setenv("DYLD_LIBRARY_PATH", fmt.Sprintf("%s%c%s", lp, os.PathListSeparator, path))
	} else {
		os.Setenv("DYLD_LIBRARY_PATH", path)
	}
}
