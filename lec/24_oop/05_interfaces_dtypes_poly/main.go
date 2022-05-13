package main

import (
	"fmt"
	"math"
	"strings"
)

const lecHeader string = "OOP: Interfaces - Dynamic Types and Polymorphism"

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

// creat func on interface
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

	// abstract type value
	// declare:
	var s shape
	pf("s (%T) = %#v\n", s, s) // => s (<nil>) = <nil>

	ball := circle{radius: 5.5}
	pf("ball (%T) = %#v\n", ball, ball) // => ball (main.circle) = main.circle{radius:5.5}

	// interface var can hold any type that implements this interface
	s = ball
	pf("s (%T) = %#v\n", s, s) // => s (main.circle) = main.circle{radius:5.5}

	// we can call func on this inteface var, that will use type, that was assigned
	printShape(s)

	pln(strings.Repeat("-", 10))

	room := rectangle{
		withd:  5,
		height: 10,
	}

	s = room
	pf("s (%T) = %#v\n", s, s)
}
