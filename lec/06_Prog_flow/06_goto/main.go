package main

import (
	"fmt"
)

func main() {
	fmt.Println("GoTo!")

	// making for loop via labels
	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
}
