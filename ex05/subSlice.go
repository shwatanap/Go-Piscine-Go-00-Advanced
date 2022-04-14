package main

import "fmt"

func subSlice(slice []int, begin int, length int, capacity int) []int {
	cap := capacity
	if capacity < length {
		cap = length
	}
	newSlice := make([]int, length, cap)
	slice_len := len(slice)
	if slice_len > cap {
		slice_len = cap
	}
	for i := 0; i+begin < slice_len; i++ {
		newSlice[i] = slice[i+begin]
	}
	return newSlice
}

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int
	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}
