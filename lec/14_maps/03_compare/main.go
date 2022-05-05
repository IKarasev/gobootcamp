package main

import (
	"fmt"
	"strings"
)

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("Maps compare!\n")
	pln(strings.Repeat("-", 10))

	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}

	//pln(a == b) // => invalid operation: a == b (map can only be compared to nil)

	// compare by string repr
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	pf("s1: %q\n", s1)
	pf("s2: %q\n", s2)

	// now we can compare s1 ad s2
	if s1 == s2 {
		pln("Maps are equal!")
	} else {
		pln("Maps are different!")
	}

	pln(strings.Repeat("-", 10))

	a = map[string]string{"A": "X"}
	b = map[string]string{"A": "X"}

	s1 = fmt.Sprintf("%s", a)
	s2 = fmt.Sprintf("%s", b)
	pf("s1: %q\n", s1)
	pf("s2: %q\n", s2)

	if s1 == s2 {
		pln("Maps are equal!")
	} else {
		pln("Maps are different!")
	}
}
