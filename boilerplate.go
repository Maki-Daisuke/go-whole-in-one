package uum7

import (
	"crypto/md5"
	"fmt"
	"os"
)

func WriteMainGo(name, version string) error {
	out, err := os.OpenFile("main.go", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = fmt.Fprintf(out, `// This file is generated by uum7-init.
//go:generate uum7 generate
package main

import(
	"os"
	"github.com/Maki-Daisuke/uum7"
)

func init() {
	uum7.Name = %q
	uum7.Version = %q
}

func main() {
	// Customize here as you like.
	uum7.Exec(os.Args[1:])
}
`, name, version)
	return err
}

func WritePackGo(data []byte) error {
	out, err := os.OpenFile("pack.go", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = fmt.Fprintf(out, `// DO NOT EDIT! This file is generated by uum7-init.
package main

import(
	//"archive/tar"
	//"compress/gzip"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"github.com/Maki-Daisuke/uum7"
)

var(
	data = %q
	hash = "%X"
)

func init() {
	unpackPath := filepath.Join(os.TempDir(), fmt.Sprintf("uum7cache-%%s-%%s-%%s", uum7.Name, uum7.Version, hash))
	err := os.Mkdir(unpackPath, 0755)
	if os.IsExist(err) {
		// Package is already unpacked
		return
	} else if err != nil {
		panic(err)
	}
	if data == "" {
		return
	}
	cmd := exec.Command("tar", "-C", unpackPath, "-zxf", "-")
	cmd.Stdin = strings.NewReader(data)
	cmd.Stdout = nil
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	os.Setenv("PATH", fmt.Sprintf("%%s%%c%%s", unpackPath, os.PathListSeparator, os.Getenv("PATH")))
}
`, data, md5.Sum(data))
	return err
}

func WritePackingList(name string) error {
	out, err := os.OpenFile("packing-list", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = fmt.Fprintf(out, `# This file is generated by uum7-init.
%s-*
`, name)
	return err
}
