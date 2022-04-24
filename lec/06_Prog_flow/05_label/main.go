package main

import (
	"fmt"
)

func main() {
	fmt.Println("Labes statement!")

	letters := [5]string{"a", "b", "c", "d", "e"}
	pairs := [2]string{"b", "c"}

	// define label
outer:
	for index, name := range letters {
		for _, pair := range pairs {
			if name == pair {
				fmt.Printf("Found pair %q at index %d\n", pair, index)
				break outer
			}
		}
	}

	outer := 1
	_ = outer //no conflict with label, as it's in another namespace
	fmt.Println("After break")
}
