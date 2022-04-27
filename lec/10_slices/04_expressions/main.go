package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slices expressions!")

	fmt.Println(strings.Repeat("-", 10))

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a: ", a)
	fmt.Printf("%T\n", a)

	// a[start:stop] - returns the slice that is part of array; start - included, stop - not
	b := a[0:3]
	fmt.Println("b := a[0:3] => ", b)
	fmt.Printf("%T\n", b)

	fmt.Println(strings.Repeat("-", 10))

	fmt.Println("a[2:] =", a[2:]) // from index to end (index included)
	fmt.Println("a[:3] =", a[:3]) // from start to index (index excluded)
	fmt.Println("a[:] =", a[:])   // all elems

	s1 := a[:] // s1 now slice of a elements
	fmt.Println("s1 = ", s1)
	fmt.Printf("%T\n", s1)

	// combine append and part
	fmt.Println(strings.Repeat("-", 10))
	s1 = append(s1[:3], 100)
	fmt.Println("s1 = append(s1[:3], 100) =>", s1)
	s1 = append(s1[:3], 200)
	fmt.Println("s1 = append(s1[:3], 200) =>", s1)

}
