package main

import "fmt"

func main() {
	fmt.Println("Const and IOTA!!")
	// consts
	const PI float64 = 3.14 // typed
	const ECON = 2.718      // untyped
	const LENGTH int = len("3.14")
	const a float64 = PI * ECON
	const b string = "I" + "con"
	const d = 5 > 7

	fmt.Println(PI, ECON, LENGTH)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)

	const x = 5
	const y float64 = 2.2 * 5
	fmt.Println(x, y)

	const r = 5
	rr := r
	fmt.Printf("r: %T \n", r)
	fmt.Printf("rr: %T \n", rr)

	// iota
	const (
		c1 = iota // starts with 0
		c2 = iota // 1
		c3 = iota // 2
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota // starts again with 0
		c22        // auto continue iota
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		North = iota
		East
		South
		West // = 3
	)
	fmt.Println(West)

	// iota step can be modified by math operations
	// e.g. by 2:
	const (
		b1 = iota * 2 // 0
		b2            // 2
		b3            // 4
	)
	fmt.Println(b1, b2, b3)

	// _ can be used to skip iota step

	const (
		d1 = -(iota + 2) // -2
		_                // skip
		d2               // -4
		d3               // -5
	)
	fmt.Println(d1, d2, d3)

}
