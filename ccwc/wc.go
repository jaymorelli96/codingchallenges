package main

import (
	"fmt"
	"io"
)

type options struct {
	c bool
}

func WordCount(r io.Reader, w io.Writer, opts options) (int, error) {
	b := make([]byte, 1024)
	n, err := r.Read(b)
	if err != nil {
		if err != io.EOF {
			return 0, err
		}
		return 0, nil
	}

	if opts.c {
		err = writeBytes(w, n)
	}

	return n, err
}

func writeBytes(w io.Writer, n int) error {
	_, err := fmt.Fprint(w, n)
	return err
}
