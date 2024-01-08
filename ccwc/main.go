package main

import (
	"flag"
)

var (
	c = flag.Bool("c", false, "print the number of bytes in the file")
)

func main() {
	flag.Parse()

}
