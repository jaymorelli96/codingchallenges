package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	c = flag.Bool("c", false, "print the number of bytes in the file")
	l = flag.Bool("l", false, "print the number of lines in the file")
	w = flag.Bool("w", false, "print the number of words in the file")
	m = flag.Bool("m", false, "print the number of characters in the file")
)

func main() {
	flag.Parse()

	args := flag.Args()

	var config options

	isDefault := !*c && !*l && !*w && !*m
	if isDefault {
		config = options{c: true, l: true, w: true}
	} else {
		config = options{c: *c, l: *l, w: *w, m: *m}
	}

	if len(args) == 0 {
		err := WordCount(os.Stdin, os.Stdout, config)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println()
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	err = WordCount(file, os.Stdout, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(os.Stdout, " %s\n", args[0])
}
