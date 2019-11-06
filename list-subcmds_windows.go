package wio

import "path/filepath"

func lookupExternalSubcmds() []string {
	exes := lookupExecutables(Name + "-*")
	for i, v := range exes {
		base := filepath.Base(v)
		ext := filepath.Ext(v)
		exes[i] = base[0 : len(base)-len(ext)]
	}
	return exes
}
