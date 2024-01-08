package main

import (
	"bytes"
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
		want string
	}

	tt := []test{
		{"file1", "3"},
		{"file2", "0"},
		{"file3", "5"},
	}

	for _, tc := range tt {
		file, err := files.Open(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		opts := options{c: true}

		var b bytes.Buffer

		err = WordCount(file, &b, opts)
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != tc.want {
			t.Errorf("got %s; want %s", b.String(), tc.want)
		}
	}
}

func TestWordCountNumberOfLines(t *testing.T) {
	files := fstest.MapFS{
		"file1": {Data: []byte("abc\ndef")},
		"file2": {Data: []byte("")},
		"file3": {Data: []byte("one line")},
		"file4": {Data: []byte("joão\n\n")},
	}

	type test struct {
		file string
		want string
	}

	tt := []test{
		{"file1", "2"},
		{"file2", "0"},
		{"file3", "1"},
		{"file4", "3"},
	}

	for _, tc := range tt {
		file, err := files.Open(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		opts := options{l: true}
		var b bytes.Buffer

		err = WordCount(file, &b, opts)
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != tc.want {
			t.Errorf("got %s; want %s", b.String(), tc.want)
		}
	}
}

func TestWordCountNumberOfWords(t *testing.T) {
	files := fstest.MapFS{
		"file1": {Data: []byte("abc\ndef\n")},
		"file2": {Data: []byte("\n")},
		"file3": {Data: []byte("thisisonebigword")},
		"file4": {Data: []byte("five words and two lines\n\n")},
		"file5": {Data: []byte("abc def")},
	}

	type test struct {
		file string
		want string
	}

	tt := []test{
		{"file1", "2"},
		{"file2", "0"},
		{"file3", "1"},
		{"file4", "5"},
		{"file5", "2"},
	}

	for _, tc := range tt {
		file, err := files.Open(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		opts := options{w: true}
		var b bytes.Buffer

		err = WordCount(file, &b, opts)
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != tc.want {
			t.Errorf("got %s; want %s", b.String(), tc.want)
		}
	}
}
