package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(n int) int {
	return n * (n + 1) / 2
}

func print_star(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if j == i-1 {
				fmt.Println("*")
			} else {
				fmt.Print("*")
			}
		}
	}
}

func main() {
	args := os.Args
	num, err := strconv.Atoi(args[1])
	if err != nil {
		os.Exit(1)
	}
	i := 1
	for sum(i) <= num {
		i++
	}
	print_star(i)
}
