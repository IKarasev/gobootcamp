package main

import (
	"fmt"
)

func main() {
	fmt.Println("For loop!")

	// no while loop in GO!
	for i := 1; i < 10; i++ {
		fmt.Printf("Count: %v\n", i)
	}

	// while example
	count := 1
	for count < 6 {
		fmt.Printf("Count %v is less then 6\n", count)
		count++
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Even num: ", i)
	}

	// break
	count = 0
	for {
		if count%7 == 0 {
			fmt.Println("div by 7: ", count)
		}
		count++
		if count == 30 {
			break
		}
	}

	fmt.Println("loops finished")
}
