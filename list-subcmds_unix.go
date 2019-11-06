// +build darwin !android,linux

package wio

import "path/filepath"

func lookupExternalSubcmds() []string {
	exes := lookupExecutables(Name + "-*")
	for i, v := range exes {
		exes[i] = filepath.Base(v)
	}
	return exes
}
