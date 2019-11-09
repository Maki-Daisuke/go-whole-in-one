package main

import (
	"bytes"
	"os"

	wio "github.com/Maki-Daisuke/go-whole-in-one"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Codec      string
	Gzip       func() `short:"z" long:"gzip" description:"Compress embedded files with Gzip (default)"`
	Xz         func() `short:"J" long:"xz" description:"Compress embedded files with Xz"`
	NoCompress func() `short:"N" long:"no-compress" description:"Do not compress embedded files"`
}

func init() {
	opts.Codec = "gzip"
	opts.Gzip = func() {
		opts.Codec = "gzip"
	}
	opts.Xz = func() {
		opts.Codec = "xz"
	}
	opts.NoCompress = func() {
		opts.Codec = ""
	}
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	files := parsePackingList()
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	err = pack(buf, files, opts.Codec)
	mustNotError(err, "")
	err = wio.WritePackGo(buf.Bytes(), opts.Codec)
	mustNotError(err, "")
}
