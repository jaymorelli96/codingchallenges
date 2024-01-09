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
		"file4": {Data: []byte("abc\ndef")},
	}

	type test struct {
		file string
		want string
	}

	tt := []test{
		{"file1", "3 "},
		{"file2", "0 "},
		{"file3", "5 "},
		{"file4", "7 "},
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
			t.Errorf("file: %s - got %s; want %s", file, b.String(), tc.want)
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
		{"file1", "1 "},
		{"file2", "0 "},
		{"file3", "0 "},
		{"file4", "2 "},
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
			t.Errorf("file: %s - got %s; want %s", file, b.String(), tc.want)
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
		{"file1", "2 "},
		{"file2", "0 "},
		{"file3", "1 "},
		{"file4", "5 "},
		{"file5", "2 "},
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
			t.Errorf("file: %s - got %s; want %s", file, b.String(), tc.want)
		}
	}
}

func TestWordCountNumberOfChars(t *testing.T) {
	files := fstest.MapFS{
		"file1": {Data: []byte("abc\ndef\n")},
		"file2": {Data: []byte("\n")},
		"file3": {Data: []byte("bigword\n")},
		"file4": {Data: []byte("joão\n")},
		"file5": {Data: []byte("abc def\n")},
		"file6": {Data: []byte("")},
	}

	type test struct {
		file string
		want string
	}

	tt := []test{
		{"file1", "8 "},
		{"file2", "1 "},
		{"file3", "8 "},
		{"file4", "5 "},
		{"file5", "8 "},
		{"file6", "0 "},
	}

	for _, tc := range tt {
		file, err := files.Open(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		opts := options{m: true}
		var b bytes.Buffer

		err = WordCount(file, &b, opts)
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != tc.want {
			t.Errorf("file: %s - got %s; want %s", file, b.String(), tc.want)
		}
	}
}

func TestWordCountDefault(t *testing.T) {
	files := fstest.MapFS{
		"file1": {Data: []byte("abc\ndef\n")},
		"file2": {Data: []byte("\n")},
		"file3": {Data: []byte("bigword\n")},
		"file4": {Data: []byte("joão\n")},
		"file5": {Data: []byte("abc def\n")},
		"file6": {Data: []byte("")},
		"file7": {Data: []byte("abc def")},
	}

	type test struct {
		file string
		want string
	}

	tt := []test{
		{"file1", "2 2 8 "},
		{"file2", "1 0 1 "},
		{"file3", "1 1 8 "},
		{"file4", "1 1 6 "},
		{"file5", "1 2 8 "},
		{"file6", "0 0 0 "},
		{"file7", "0 2 7 "},
	}

	for _, tc := range tt {
		file, err := files.Open(tc.file)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		opts := options{
			c: true,
			l: true,
			w: true,
		}

		var b bytes.Buffer

		err = WordCount(file, &b, opts)
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != tc.want {
			t.Errorf("file: %s - got %s; want %s", file, b.String(), tc.want)
		}
	}
}
