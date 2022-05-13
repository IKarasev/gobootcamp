package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "OOP: Empty Interface"

// empty interface - interface with no methods
// such interface is implemented by any type
// so it can contain any value!

type emptyInterface interface {
}

type pet struct {
	name   string
	hungry bool
}

// we can use empty interfaces in struct to create field of any type
type person struct {
	info interface{}
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	var empty interface{}
	pln("> empty:")
	pf("type = %T\nvalue = %#v\n\n", empty, empty)

	empty = 5
	pln("> empty = 5:")
	pf("type = %T\nvalue = %#v\n\n", empty, empty)

	empty = "Some"
	pln(`> empty = "Some":`)
	pf("type = %T\nvalue = %#v\n\n", empty, empty)

	empty = pet{name: "Charly", hungry: true}
	pln(`> empty = pet{name: "Charly", hungry: true}:`)
	pf("type = %T\nvalue = %#v\n\n", empty, empty)

	// try slice
	empty = []float64{1.1, 2.2, 3.3, 4.4}
	pln(`> empty = []float64{1.1, 2.2, 3.3, 4.4}:`)
	pf("type = %T\nvalue = %#v\n\n", empty, empty)

	// get len of slice
	// pln("len(empty) =", len(empty)) => error: invalid argument: empty (variable of type interface{}) for len
	// use assertion:
	sl, ok := empty.([]float64)

	if ok {
		pln("len(empty) =", len(sl))
	}

	pln(strings.Repeat("-", 10))

	// try to work with person struct
	you := person{}
	pln(`> you := person{}`)
	pln("info: ", you.info)

	you.info = "John Malkovich"
	pln(`> you.info = "John Malkovich"`)
	pln("info: ", you.info)

	you.info = 50
	pln(`> you.info = 50`)
	pln("info: ", you.info)

	you.info = []float64{1.5, 5.4, 10.}
	pln(`> you.info = []float64{1.5,5.4,10.}`)
	pln("info: ", you.info)

}
