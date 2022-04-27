package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slices copy!")

	fmt.Println(strings.Repeat("-", 10))

	nums := []int{1, 2}
	fmt.Println(nums)

	// append
	nums = append(nums, 3) // returns new slice!
	fmt.Println(nums)

	// append multiple items
	nums = append(nums, 4, 5, 6)
	fmt.Println(nums)

	// append one slice to another
	n := []int{10, 20}
	nums = append(nums, n...)
	fmt.Println(nums)

	fmt.Println(strings.Repeat("-", 10))

	// copy slices
	src := []int{100, 200, 300}
	dst := make([]int, len(src)) // creates the new slice of src length
	nn := copy(dst, src)         // copies slice and returns # of elems copied
	fmt.Println(src, dst, nn)

	fmt.Println(strings.Repeat("-", 10))

	// copy slice to smaller one
	src = []int{100, 200, 300}
	dst = make([]int, 2) // creates the new slice of src length
	nn = copy(dst, src)  // copies slice and returns # of elems copied
	fmt.Println(src, dst, nn)

	fmt.Println(strings.Repeat("-", 10))

	// copy slice to zero slice
	src = []int{100, 200, 300}
	dst = make([]int, 0) // creates the new slice of src length
	nn = copy(dst, src)  // copies slice and returns # of elems copied
	fmt.Println(src, dst, nn)
}
