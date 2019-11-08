package wio

import (
	"fmt"
	"os"
)

// SetEnv is called in by pack.go during preparation of environment.
// You should not call this.
func SetEnv(path string) {
	os.Setenv("WIOPATH", path)
	os.Setenv("PATH", fmt.Sprintf("%s%c%s", path, os.PathListSeparator, os.Getenv("PATH")))
	setEnv(path)
}
