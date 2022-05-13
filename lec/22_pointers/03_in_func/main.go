package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Pointers: Passing and returning in functions"

// function takes pointer as arg and returns pointer
func f1(n *int) *float64 {
	fmt.Println("> f1()")
	*n = 100 // changes value of original var

	b := 8.43
	return &b
}

// without pointer
func f2(n int) {
	fmt.Println("> f2()")
	n = 999
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	x := 18
	p := &x

	pln("x =", x)

	// call f1 to change x
	r := f1(p)
	pln("x =", x)
	pf("r type %T = %v, origin val = %v\n", r, r, *r)

	// call f2 to change x
	f2(x)
	pln("x =", x) // var x not changed, functions can't change vars passed to it
	// as functions work with copies of var
	// by passing pointer, function can change value of its var
}
