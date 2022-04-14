package main

import "fmt"

type FortyTwo struct{}

func main() {
	vars := []interface{}{"42", uint(42), 42, byte(42), int16(42), uint32(42), int64(42), 42 != 42, float32(42), 42.0, complex64(42 + 0i), (42 + 21i), FortyTwo{}, [1]int{42}, map[string]int{"42": 42}, (*int)(nil), []int(nil), (chan int)(nil), nil}
	for _, v := range vars {
		switch v.(type) {
		case int, int16, int64:
			fmt.Printf("%v is an %T.\n", v, v)
		case *int:
			fmt.Printf("%p is an %T.\n", v, v)
		default:
			fmt.Printf("%v is a %T.\n", v, v)
		}
	}
}
