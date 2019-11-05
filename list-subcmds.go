package wio

import (
	"fmt"
	"sort"
	"strings"
)

func ListSubcommands() ([]string, error) {
	found := make([]string, 0, len(builtins))
	for name := range builtins {
		if !strings.HasPrefix(name, "-") {
			found = append(found, name)
		}
	}
	cmds, err := LookupCmdNames(fmt.Sprintf("%s-*", Name))
	if err != nil {
		return nil, err
	}
	for i, c := range cmds {
		cmds[i] = c[len(Name)+1:]
	}
	found = append(found, cmds...)
	sort.Strings(found)
	return found, nil
}
