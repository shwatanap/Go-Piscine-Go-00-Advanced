package main

import (
	"flag"
	"fmt"
	"regexp"
	"unicode/utf8"
)

func main() {
	var email_pattern = `^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`
	var email_re = regexp.MustCompile(email_pattern)

	flag.Parse()
	args := flag.Args()
	size := len(args)
	if size == 0 {
		fmt.Println("Invalid argument")
	} else {
		for i := 0; i < size; i++ {
			// fmt.Println(len(email_re.FindAllString(args[i], -1)) != 0)
			is_valid_email_format := (len(email_re.FindAllString(args[i], -1)) != 0)
			if utf8.RuneCountInString(args[i]) <= 256 && is_valid_email_format {
				fmt.Println(args[i], "is a valid email address.")
			} else {
				fmt.Println(args[i], "is NOT a valid email address.")
			}
		}
	}
}
