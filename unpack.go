package wio

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Unpack is called in Deployment phase by pack.go.
// You should not call this.
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
	// For security consideration, all unpacked files can be accessed only by the command user.
	// Because, we don't know whether sensitive/confidential data is embedded in the binary.
	// However, anyone can run the command since each user has her/his own cache directory.
	// BTW, super users (e.g. root user) can access to unpacked files anyhow.
	perm := hdr.Mode & 0700
	file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(perm))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := io.Copy(file, rd); err != nil {
		panic(err)
	}
}
