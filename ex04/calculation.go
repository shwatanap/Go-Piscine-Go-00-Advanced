package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(args []string) (string, bool) {
	if len(args) != 3 {
		return "", false
	} else {
		num1, err := strconv.Atoi(args[1])
		if err != nil {
			return "", false
		}
		num2, err := strconv.Atoi(args[2])
		if err != nil || num2 == 0 {
			return "", false
		}
		sum := num1 + num2
		difference := num1 - num2
		product := num1 * num2
		quotient := num1 / num2
		s := fmt.Sprintf("sum: %d\ndifference: %d\nproduct: %d\nquotient: %d\n", sum, difference, product, quotient)
		return s, true
	}
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
