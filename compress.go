package wio

import (
	"compress/gzip"
	"io"

	"github.com/ulikunitz/xz"
)

// CompressWriter compresses w using codec.
func CompressWriter(w io.Writer, codec string) io.Writer {
	switch codec {
	case "gzip":
		return gzip.NewWriter(w)
	case "xz":
		w, err := xz.NewWriter(w)
		if err != nil {
			panic(err)
		}
		return w
	default:
		return w
	}
}

// uncompressReader uncompresses r using codec.
func uncompressReader(r io.Reader, codec string) io.Reader {
	switch codec {
	case "gzip":
		r, err := gzip.NewReader(r)
		if err != nil {
			panic(err)
		}
		return r
	case "xz":
		r, err := xz.NewReader(r)
		if err != nil {
			panic(err)
		}
		return r
	default:
		return r
	}
}
