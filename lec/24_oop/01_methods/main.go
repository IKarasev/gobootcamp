package main

import (
	"fmt"
	"strings"
	"time"
)

const lecHeader string = "OOP: methods"

// go doe not have classic classes
// but methods can be applied in named types

// create new type
type names []string

// creating a method for new type
func (n names) print() {
	// n - method reciever, can be viewd as "self" or "this" in other lang
	for i, v := range n {
		fmt.Println(i, ":", v)
	}
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	const day = 24 * time.Hour

	pf("day ~ %T\n", day) // => time.Duration
	// its named type that has methods

	seconds := day.Seconds()
	pf("seconds(%T) = %v\n", seconds, seconds) // => float64 and 86400

	// using names type and its method
	friends := names{"Andrei", "Mike"}
	friends.print()

	// another way to call method from type
	names.print(friends)

	// method is a function that takes named type variable as param
	// method belongs to type namesapce
	// function - to package namespace

	// converting var types
	var n int64 = 4848848
	pln("n =", n)
	pln(time.Duration(n))

	// its possible only if underlying type is the same (time underlying type is int64)
}
