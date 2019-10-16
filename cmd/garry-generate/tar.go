package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func pack(out io.Writer, packList []string) error {
	gz := gzip.NewWriter(out)
	defer gz.Close()
	tw := tar.NewWriter(gz)
	defer tw.Close()
	for _, path := range packList {
		path, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		err = walkAndWriteToTar(tw, filepath.Base(path), path)
		if err != nil {
			return err
		}
	}
	return nil
}

func walkAndWriteToTar(tw *tar.Writer, tarPath string, fsPath string) error {
	file, err := os.Open(fsPath)
	if err != nil {
		return err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	if fi.Mode()&os.ModeSymlink != 0 { // Follow symlink
		p, err := filepath.EvalSymlinks(fsPath)
		if err != nil {
			return err
		}
		return walkAndWriteToTar(tw, tarPath, p)
	}
	if fi.IsDir() {
		return writeDirAndWalk(tw, tarPath, fsPath, file)
	}
	if fi.Mode().IsRegular() {
		return writeFile(tw, tarPath, file)
	}
	// Ignore all other non-regular files
	return nil
}

func writeFile(tw *tar.Writer, tarPath string, file *os.File) error {
	fi, _ := file.Stat() // must not be error
	h, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		return err
	}
	h.Name = tarPath
	err = tw.WriteHeader(h)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}

func writeDirAndWalk(tw *tar.Writer, tarPath string, fsPath string, file *os.File) error {
	fi, _ := file.Stat() // must not be error
	h, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		return err
	}
	h.Name = fmt.Sprintf("%s%c", tarPath, '/') // Add trailing slash to indicate a directory
	err = tw.WriteHeader(h)
	if err != nil {
		return err
	}
	names, err := file.Readdirnames(0) // Read all dirnames
	if err != nil {
		return err
	}
	for _, n := range names {
		if n == "." || n == ".." {
			continue
		}
		err := walkAndWriteToTar(tw, fmt.Sprintf("%s/%s", tarPath, n), filepath.Join(fsPath, n))
		if err != nil {
			return err
		}
	}
	return nil
}
