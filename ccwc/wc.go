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
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		if err != io.EOF {
			return err
		}

		handleEmptyFile(w, opts)

		return nil
	}

	b := buf.Bytes()

	if opts.l {
		count := bytes.Count(b, []byte{'\n'})
		fmt.Fprintf(w, "%d ", count)
	}

	if opts.w {
		f := bytes.Fields(b)
		fmt.Fprintf(w, "%d ", len(f))
	}

	if opts.c {
		fmt.Fprintf(w, "%d ", len(b))
	}

	if opts.m {
		n := utf8.RuneCount(b)
		fmt.Fprintf(w, "%d ", n)
	}

	return err
}

func handleEmptyFile(w io.Writer, opts options) {
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
}
