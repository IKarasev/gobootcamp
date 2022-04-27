package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slices operations in-depth!")
	fmt.Println(strings.Repeat("-", 10))

	var nums []int // nill, no backing array
	fmt.Printf("nums = %#v\n", nums)
	fmt.Printf("len = %d\n", len(nums)) // => 0
	fmt.Printf("cap = %d\n", cap(nums)) // => 0

	nums = append(nums, 1, 2)
	fmt.Printf("nums = %#v\n", nums)
	fmt.Printf("len = %d\n", len(nums)) // => 2
	fmt.Printf("cap = %d\n", cap(nums)) // => 2

	nums = append(nums, 3)
	fmt.Printf("nums = %#v\n", nums)
	fmt.Printf("len = %d\n", len(nums)) // => 3
	fmt.Printf("cap = %d\n", cap(nums)) // => 4, go added 1 extra spot in backing array
	// on append go creates new backing array if cap of new array is > then origin cap
	// go adds extra spots to backing array in order to escape creating new backing array
	// on next append

	nums = append(nums, 4)
	fmt.Printf("nums = %#v\n", nums)
	fmt.Printf("len = %d\n", len(nums)) // => 4
	fmt.Printf("cap = %d\n", cap(nums)) // => 4, go didn't create new backed array, but added new val to existing one

	nums = append(nums, 5)
	fmt.Printf("nums = %#v\n", nums)
	fmt.Printf("len = %d\n", len(nums)) // => 5
	fmt.Printf("cap = %d\n", cap(nums)) // => 8, 3 extra

	nums = append(nums[:4], 100, 200, 300, 400, 500)
	fmt.Printf("nums = %#v\n", nums)
	fmt.Printf("len = %d\n", len(nums)) // => 8
	fmt.Printf("cap = %d\n", cap(nums)) // => 16, 8 extra

	fmt.Println(strings.Repeat("-", 10))

	letters := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Printf("letters = %#v\n", letters)
	fmt.Printf("len = %d\n", len(letters)) // => 6
	fmt.Printf("cap = %d\n", cap(letters)) // => 6

	letters = append(letters[:1], "x", "y")
	fmt.Printf("letters = %#v\n", letters)
	fmt.Printf("len = %d\n", len(letters)) // => 3
	fmt.Printf("cap = %d\n", cap(letters)) // => 6, cap didn't change

	//fmt.Println(letters[4]) // => panic: runtime error: index out of range [4] with length 3
	// but we can make slice operation till backed array size
	fmt.Println(letters[3:6]) // => [d e f], taken from backed array

	fmt.Println(strings.Repeat("-", 10))
	fmt.Println("Slice backed array increment")

	var s []int
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("len = %d\n", len(s)) // => 0
	fmt.Printf("cap = %d\n", cap(s)) // => 0

	s = append(s, 1)
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("len = %d\n", len(s)) // => 1
	fmt.Printf("cap = %d\n", cap(s)) // => 1

	s = append(s, 2)
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("len = %d\n", len(s)) // => 2
	fmt.Printf("cap = %d\n", cap(s)) // => 2

	s = append(s, 3)
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("len = %d\n", len(s)) // => 3
	fmt.Printf("cap = %d\n", cap(s)) // => 4, +1

	s = append(s, 4, 5)
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("len = %d\n", len(s)) // => 5
	fmt.Printf("cap = %d\n", cap(s)) // => 8, +3

	s = append(s, 6, 7, 8, 9)
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("len = %d\n", len(s)) // => 9
	fmt.Printf("cap = %d\n", cap(s)) // => 16, +7

	s = append(s, 10, 11, 12, 13, 14, 15, 16, 17)
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("len = %d\n", len(s)) // => 17
	fmt.Printf("cap = %d\n", cap(s)) // => 32, +15

}
