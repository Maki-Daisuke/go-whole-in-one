package wio

import (
	"path/filepath"
)

func LookupCmdNames(pattern string) ([]string, error) {
	paths, err := LookupCmds(pattern)
	if err != nil {
		return nil, err
	}
	for i, p := range paths {
		paths[i] = filepath.Base(p)
	}
	return paths, nil
}

func LookupCmds(pattern string) (cmds []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	return lookupCmds(pattern), nil
}

func IsExecutable(path string) (ok bool, err error) {
	defer func() {
		if e := recover(); e != nil {
			ok = false
			err = e.(error)
		}
	}()
	return isExecutable(path), nil
}
