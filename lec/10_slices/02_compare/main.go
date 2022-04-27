package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slices compare!")

	// declared but not init
	var n []int
	fmt.Println("n: ", n, " is nill: ", n == nil)

	// empty init slice
	m := []int{}
	fmt.Println("m: ", m, " is nill: ", m == nil)

	fmt.Println(strings.Repeat("-", 10))

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	//fmt.Println(a == b) //=> invalid operation: a == b (slice can only be compared to nil)
	// slices could be compared only elem by elem
	eq := true
	for i, va := range a {
		if va != b[i] {
			eq = false
			break
		}
	}
	fmt.Println("a == b: ", eq)

	fmt.Println(strings.Repeat("-", 10))

	a, b = []int{1, 2, 3}, []int{1, 5, 3}
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	eq = true
	for i, va := range a {
		if va != b[i] {
			eq = false
			break
		}
	}
	fmt.Println("a == b: ", eq)

	// note! check for len of slices!

	fmt.Println(strings.Repeat("-", 10))

	a, b = []int{1, 2, 3}, []int{1, 2, 3}
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	eq = true

	if len(a) == len(b) {
		for i, va := range a {
			if va != b[i] {
				eq = false
				break
			}
		}
	} else {
		eq = false
	}
	fmt.Println("a == b: ", eq)

	fmt.Println(strings.Repeat("-", 10))

	a, b = nil, []int{1, 2, 3}
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	eq = true

	if len(a) == len(b) {
		for i, va := range a {
			if va != b[i] {
				eq = false
				break
			}
		}
	} else {
		eq = false
	}
	fmt.Println("a == b: ", eq)
}
