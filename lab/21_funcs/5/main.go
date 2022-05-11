package main

import (
	"fmt"
)

func f1(n ...int) (s int) {
	for _, v := range n {
		s += v
	}

	return
}

func main() {
	fmt.Println("f1(1,2,3,4) =", f1(1, 2, 3, 4))
}
