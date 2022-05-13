package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Pointers: Passing value vs Passing pointer"

func changeVals(i int, f float64, s string, b bool) {
	fmt.Println("> changeVals()")
	i = 99
	f = 3.14
	s = "New string"
	b = false
}

func changeValsByPointer(i *int, f *float64, s *string, b *bool) {
	fmt.Println("> changeValsByPointer()")
	*i = 99
	*f = 3.14
	*s = "New string"
	*b = false
}

// POINTERS TO STRUCT
// struct to use in func
type Product struct {
	price float64
	name  string
}

func changeProduct(p Product) {
	fmt.Println("> changeProduct()")
	p.price = 100.
	p.name = "NoName"
}

// Pointer to struct
func changeProductByPointer(p *Product) {
	fmt.Println("> changeProductByPointer()")
	(*p).price = 100.
	(*p).name = "NoName"
}

// SLICES
func changeSlice(n []int) {
	fmt.Println("> changeSlice()")
	for i := 0; i < len(n); i++ {
		n[i] = 22
	}
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	i, f, s, b := 5, 100.0, "Something", true
	pln("Befor change")
	pln(i, f, s, b)

	changeVals(i, f, s, b)
	pln("After change")
	pln(i, f, s, b) // nothing changed

	changeValsByPointer(&i, &f, &s, &b)
	pln("After change")
	pln(i, f, s, b)

	pln(strings.Repeat("-", 10))

	p := Product{
		price: 10.15,
		name:  "T-shirt",
	}
	pln("p = ", p)

	changeProduct(p)
	pln("p = ", p) // same, not changed

	changeProductByPointer(&p)
	pln("p = ", p)

	pln(strings.Repeat("-", 10))

	sl := []int{1, 2, 3, 4}
	pln("sl =", sl)
	changeSlice(sl)
	pln("sl =", sl) // slice has changed!!!
	// slice variable stores the pointer to backing array, not the array itself
	// same goes with maps
}
