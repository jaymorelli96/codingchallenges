package main

import (
	"fmt"
	"io"
)

type options struct {
	c bool
	l bool
}

func WordCount(r io.Reader, w io.Writer, opts options) error {
	b := make([]byte, 1024)
	n, err := r.Read(b)
	if err != nil {
		if err != io.EOF {
			return err
		}

		return write(w, 0)
	}

	if opts.c {
		err = write(w, n)
	}

	if opts.l {
		count := 1
		for _, v := range b {
			if v == '\n' {
				count++
			}
		}

		err = write(w, count)
	}

	return err
}

func write(w io.Writer, n int) error {
	_, err := fmt.Fprint(w, n)
	return err
}
