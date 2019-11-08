// +build !android,linux

package wio

import (
	"fmt"
	"os"
)

func setEnv(path string) {
	if llp := os.Getenv("LD_LIBRARY_PATH"); llp != "" {
		os.Setenv("LD_LIBRARY_PATH", fmt.Sprintf("%s%c%s", llp, os.PathListSeparator, path))
	} else {
		os.Setenv("LD_LIBRARY_PATH", path)
	}
}
