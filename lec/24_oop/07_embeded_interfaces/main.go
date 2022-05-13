package main

import (
	"fmt"
	"math"
	"strings"
)

const lecHeader string = "OOP: Embeded Interfaces"

// in go inheritance is not supported!
// so interfaces can't extend or inherit from other interfaces

// but we can merge two interfaces in one
// it's called interface embedding

// two interfaces with different methods
type shape interface {
	area() float64
}

type obj interface {
	volume() float64
}

// now we can create another interface, that includes previous two (embedding)

type geometry interface {
	shape             // embed shape
	obj               // embed obj
	getColor() string // new method
}

// NOTE: circular embeding is not allowed!

// define types that implement geometry
type cube struct {
	edge float64
}

// area() from shape interface
func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

// volume() from obj interface
func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

// getColor() from geometry interface
func (c cube) getColor() string {
	return "No color"
}

// define func on geometry interface
func measure(g geometry) (a float64, v float64) {
	fmt.Printf("> measure(%#v)\n", g)
	a = g.area()
	v = g.volume()
	return
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	c := cube{edge: 5.}
	a, v := measure(c)
	pf("Area = %.3f\nVolume = %.3f\n", a, v)
}
