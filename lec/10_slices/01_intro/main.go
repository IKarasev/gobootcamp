package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slices!")
	fmt.Println(strings.Repeat("-", 10))

	// init slice
	var cities []string
	fmt.Println("Is nill: ", cities == nil)
	fmt.Printf("%#v\n", cities)
	fmt.Println(len(cities))

	// cities[0] = "Toronto" // => panic: runtime error: index out of range [0] with length 0
	// can't assign not init element

	fmt.Println(strings.Repeat("-", 10))

	// init slice with items
	nums := []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)

	// init slice with 2 elemets
	n := make([]int, 2)
	fmt.Printf("%#v\n", n)

	fmt.Println(strings.Repeat("-", 10))

	// create new type of slice
	type names = []string
	friends := names{"Anton", "Andrei"}
	fmt.Printf("%#v\n", friends)

	myFriend := friends[0]
	fmt.Printf("My friend is: %s\n", myFriend)

	fmt.Println(strings.Repeat("-", 10))

	// iteration
	for i := 0; i < len(nums); i++ {
		fmt.Printf("number at %d: %d\n", i, nums[i])
	}
	for i, v := range nums {
		fmt.Printf("num at %d: %d\n", i, v)
	}

	fmt.Println(strings.Repeat("-", 10))

	// assign slices
	var nn []int
	nn = nums
	fmt.Printf("%#v\nlen: %d\n", nn, len(nn))

}
