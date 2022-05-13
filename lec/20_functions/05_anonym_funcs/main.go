package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Functions: anonymous functions"

// function that returnes another func
// also called first class functions
func incr(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// define anonim func
	func(msg string) {
		pln(msg) // can use main() vars, as it is in main namespace
	}("anonym func") // give the arg emidiatly, as func will be executed at this point

	// anonym funcs can be used to wrap other funcs

	a := incr(5)
	pf("a ~ %T\n", a)
	pln("a =", a) // mem pointer to a

	// we can call on a, as it's func
	pln("+1 =", a())
	pln("+1 =", a())
}
