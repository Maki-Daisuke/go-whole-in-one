package wio

// LookupExecutables searches all executables matching pattern in
// your PATH environment variable, and returns absolute file paths
// of the executables.
// pattern is interpreted as file glob. See path/filepath#Glob for
// details of file glob syntax.
func LookupExecutables(pattern string) (paths []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	return lookupExecutables(pattern), nil
}

// IsExecutable reports whether the file indicated by path is
// an executable  in a platform-dependent manner.
func IsExecutable(path string) (ok bool, err error) {
	defer func() {
		if e := recover(); e != nil {
			ok = false
			err = e.(error)
		}
	}()
	return isExecutable(path), nil
}
