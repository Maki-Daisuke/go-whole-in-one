package wio

func LookupExecutables(pattern string) (paths []string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	return lookupExecutables(pattern), nil
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
