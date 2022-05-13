package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Pointers: Address of and Deferencing Operators"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	name := "Coppola"
	pln("ponter to name val (mem address):", &name)

	// init pointer
	p := &name

	pf("p %T: %v\n", p, p)
	pf("Address of name is %p\n", &name)
	pln("Address of pointer: ", &p)

	// declare pointer withought val
	var p1 *float64 // it is the pointer to float64 var

	n := 7.6
	p1 = &n
	pf("New pointer type %T = %p\n", p1, p1)

	// another way to declare a pointer
	p2 := new(int) // pointer to int val
	var x int = 54
	p2 = &x
	pf("p2 type %T = %p\n", p2, p2)

	// referencing to pointer var value
	*p2 = 99 // * - deferencing operator
	pln("New x val via pointer =", x)
	pln("*p2 == x:", *p2 == x)

}
