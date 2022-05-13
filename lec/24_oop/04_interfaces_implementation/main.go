package main

import (
	"fmt"
	"math"
	"strings"
)

const lecHeader string = "OOP: Interfaces intro"

// define interface
// we only describe methods in interface
// NOTE: interfaces applied implicintly
// we do not need to declare interface for each type it implemets
// but types that implemetn interface must define all methods in that
// interface
type shape interface {
	area() float64
	perimeter() float64
}

// define new types
type rectangle struct {
	withd, height float64
}

type circle struct {
	radius float64
}

// create methods to calc area
func (r rectangle) area() float64 {
	return r.height * r.withd
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.withd)
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

// create funcs to print shape params
// instead of creating separate funcs for rectangle and circle
// make one func fo shape interface
func printShape(s shape) {
	fmt.Println("> printShape()")
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area = %.3f\n", s.area())
	fmt.Printf("Perimeter = %.3f\n", s.perimeter())
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	myRec := rectangle{
		withd:  10.,
		height: 5.,
	}

	myCircle := circle{
		radius: 10.,
	}

	pln("myRec =", myRec)
	pln("myCircle =", myCircle)

	printShape(myCircle)
	printShape(myRec)

}
