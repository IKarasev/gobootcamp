package main

import (
	"fmt"
)

func f1(n ...int) int {
	s := 0

	for _, v := range n {
		s += v
	}

	return s
}

func main() {
	fmt.Println("f1(1,2,3,4) =", f1(1, 2, 3, 4))
}
