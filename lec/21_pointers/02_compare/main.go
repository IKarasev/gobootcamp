package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Pointers: Pointer to Pinter & Compare"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	a := 5.5
	p1 := &a
	pp1 := &p1
	pln("a = ", a)
	pf("Pointer to a, p1 %T = %p\n", p1, p1)
	pf("Pointer to poiter, pp1 %T = %p\n", pp1, pp1)
	pln(strings.Repeat("-", 10))
	pf("*p = %v\n", *p1)
	pf("*pp1 = %v\n", *pp1)
	pf("**pp1 = %v\n", **pp1)

	pln(strings.Repeat("-", 10))

	// Comparing pointers
	// pointers are equal if they point to the same var
	var p2 *int
	pln("p2 = ", p2)
	pf("p2 is %#v\n", p2)

	y := 5
	z := 5

	p2 = &y
	p3 := &z
	pln("p2 == p3:", p2 == p3)

}
