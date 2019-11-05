package wio

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unpack(dest string, data io.Reader) {
	gz, err := gzip.NewReader(data)
	if err != nil {
		panic(err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			panic(err)
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			writeDir(dest, hdr)
			break
		case tar.TypeReg:
			writeFile(dest, hdr, tr)
			break
		default:
			panic(fmt.Errorf("Unknown file type: Typeflag=%d, Name=%s", hdr.Typeflag, hdr.Name))
		}
	}
}

func writeDir(dest string, hdr *tar.Header) {
	name := filepath.Join(dest, hdr.Name)
	err := os.MkdirAll(name, 0755)
	if err != nil {
		panic(err)
	}
}

func writeFile(dest string, hdr *tar.Header, rd io.Reader) {
	name := filepath.Join(dest, hdr.Name)
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		panic(err)
	}
	// Set file permission based on the original owners permission, regardless of group or others.
	// Because, we can never expect under which user-ID this file is created.
	// But whoever created, it is probably needed to be read, executed and written as the author does.
	perm := hdr.Mode & 0700
	perm |= perm >> 3
	perm |= perm >> 6
	file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(perm))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := io.Copy(file, rd); err != nil {
		panic(err)
	}
}
