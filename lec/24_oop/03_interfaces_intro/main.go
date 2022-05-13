package main

import (
	"fmt"
	"math"
	"strings"
)

const lecHeader string = "OOP: Interfaces intro"

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
func printCircle(c circle) {
	fmt.Println("> printCircle()")
	fmt.Printf("Shape: %+v\n", c)
	fmt.Printf("Area = %.3f\n", c.area())
	fmt.Printf("Perimeter = %.3f\n", c.perimeter())

}

func printRecangle(r rectangle) {
	fmt.Println("> printRectangle()")
	fmt.Printf("Shape: %+v\n", r)
	fmt.Printf("Area = %.3f\n", r.area())
	fmt.Printf("Perimeter = %.3f\n", r.perimeter())
}

// not good approach, as print funcs follows same logic for both
// rectangle and circle

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

	printCircle(myCircle)
	printRecangle(myRec)

}
