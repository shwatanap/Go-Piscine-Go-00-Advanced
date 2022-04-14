package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		n = flag.Bool("n", false, "omit trailing newline")
		s = flag.String("s", " ", "separator")
	)
	flag.Parse()
	args := flag.Args()
	argLen := len(args)
	for i := 0; i < argLen; i++ {
		fmt.Print(args[i])
		if i == argLen-1 {
			if !*n {
				fmt.Print("\n")
			}
		} else {
			fmt.Print(*s)
		}
	}
	if argLen == 0 {
		fmt.Print("\n")
	}
}
