package main

import (
	"fmt"
	"os"
)

func mustNotError(err error, msg string) {
	if err != nil {
		if msg == "" {
			msg = fmt.Sprintf("Error: %s", err)
		}
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}
