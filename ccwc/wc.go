package main

import (
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

type options struct {
	c bool
	l bool
	w bool
	m bool
}

func WordCount(r io.Reader, w io.Writer, opts options) error {
	b := make([]byte, 1024)
	n, err := r.Read(b)
	if err != nil {
		if err != io.EOF {
			return err
		}

		if opts.c {
			fmt.Fprintf(w, "0 ")
		}
		if opts.l {
			fmt.Fprintf(w, "0 ")
		}
		if opts.w {
			fmt.Fprintf(w, "0 ")
		}
		if opts.m {
			fmt.Fprintf(w, "0 ")
		}
		return nil
	}

	b = b[:n]

	if opts.c {
		fmt.Fprintf(w, "%d ", len(b))
	}

	if opts.l {
		count := 1
		for _, v := range b {
			if v == '\n' {
				count++
			}
		}

		fmt.Fprintf(w, "%d ", count)
	}

	if opts.w {
		f := bytes.Fields(b)
		fmt.Fprintf(w, "%d ", len(f))
	}

	if opts.m {
		n := utf8.RuneCount(b)
		fmt.Fprintf(w, "%d ", n)
	}

	return err
}
