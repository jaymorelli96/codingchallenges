package main

import (
	"io"
	"testing"
	"testing/fstest"
)

func TestWordCountNumberOfBytes(t *testing.T) {
	files := fstest.MapFS{
		"file1": {Data: []byte("abc")},
		"file2": {Data: []byte("")},
		"file3": {Data: []byte("joão")},
	}

	type test struct {
		file string
		want int
	}

	tt := []test{
		{"file1", 3},
		{"file2", 0},
		{"file3", 5},
	}

	for _, tc := range tt {
		file, err := files.Open(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		opts := options{c: true}

		got, err := WordCount(file, io.Discard, opts)
		if err != nil {
			t.Fatal(err)
		}

		if got != tc.want {
			t.Errorf("got %d; want %d", got, tc.want)
		}
	}
}

