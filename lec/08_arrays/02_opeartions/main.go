package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Array operations!")

	numbers := [3]int{1, 2, 3}
	fmt.Printf("%#v\n", numbers)

	// assign element
	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	// enumerated array loop
	for i, v := range numbers {
		fmt.Println("index: ", i, " value: ", v)
	}

	fmt.Println(strings.Repeat("-", 20)) // note: strings.Repeat creates string of give chars repeated n times

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index: ", i, " Value: ", numbers[i])
	}

	fmt.Println(strings.Repeat("-", 20))

	// multy dim array
	bal := [2][3]int{
		{1, 2, 3}, // two ways of init multy dim array
		[3]int{10, 20, 30},
	}
	fmt.Printf("%#v\n", bal)
	fmt.Println(bal)

	// copy array
	fmt.Println(strings.Repeat("-", 20))
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

	m[1] = 54
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

	n[0] = 10
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

	// compare array
	fmt.Println(strings.Repeat("-", 20))
	n = m
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)
	fmt.Println("m equals n: ", m == n)

	n[2] = 54
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)
	fmt.Println("m equals n: ", m == n)
}
