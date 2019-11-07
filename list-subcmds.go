package wio

import (
	"sort"
	"strings"
)

// ListSubcommands returns names of all available subcommands
// in sorted order.
// It is handy to implement your customized help message.
func ListSubcommands() (cmds []string, err error) {
	cmds = make([]string, 0, len(builtins))
	check := make(map[string]bool)

	for name := range builtins {
		if !strings.HasPrefix(name, "-") {
			check[name] = true
			cmds = append(cmds, name)
		}
	}

	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			cmds = nil
		}
	}()
	for _, c := range lookupExternalSubcmds() {
		c = c[len(Name)+1:]
		if _, exists := check[c]; exists {
			continue
		}
		check[c] = true
		cmds = append(cmds, c)
	}
	sort.Strings(cmds)
	return cmds, nil
}
