package main

import (
	"fmt"
	"math"
	"strings"
)

const lecHeader string = "OOP: Type Assertions and Type Switches"

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

// type that doesn't implement shape interface
type triangle struct {
	side1  float64
	side2  float64
	side3  float64
	angle1 float64
	angle2 float64
	angle3 float64
}

// create methods to calc area
func (r rectangle) area() float64 {
	return r.height * r.withd
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.withd)
}

func (r rectangle) diagonal() float64 {
	return math.Sqrt(math.Pow(r.height, 2) + math.Pow(r.withd, 2))
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

// circle func to calc sphere volume
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// interface func to print shape val, area, perimeter
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

	// type assertion provides access to interface concrete value
	var s shape = circle{radius: 5.5}
	pf("s ~ %T\n", s) // => s ~ main.circle

	// s.volume() // error: s.volume undefined (type shape has no field or method volume)
	// s is still interface var with dynamic type that doesn't have volume() method

	// to call volume() method from circle
	// type assirtion is used
	v := s.(circle).volume()
	pln("volume of radius 5.5 =", v)

	// type assirtion can fail, when requested dynamic is not equal to actual dynamic type
	// eg: s.(rectangle)
	// we can chek that assirtion is succesfull
	r, ok := s.(rectangle)          // r - value of dynamic type, ok - bool success value
	pf("Assirtion value: %#v\n", r) // => main.rectangle{withd:0, height:0}, zero value of requested type
	pf("Success: %v\n", ok)         // => false, s dynamic type is circle

	c, ok := s.(circle)
	pf("Assirtion value: %#v\n", c) // => main.circle{radius:5.5}, our circle val
	pf("Success: %v\n", ok)         // => true

	if ok {
		pln("Volume =", c.volume()) // can call volume() from recived var
	} else {
		pln("Failed assirtion")
	}

	pln(strings.Repeat("-", 10))

	// type switches (do something depending on type)
	switch val := s.(type) {
	case circle:
		pf("%#v is circle\n", val)
		pf("volume = %.3f\n", val.volume())
	case rectangle:
		pf("%#v is rectangle\n", val)
		pf("diagonal = %.3f\n", val.diagonal())
	}

	// change s dynamic type to rectangle and try switch again
	s = rectangle{withd: 5, height: 10}
	switch val := s.(type) {
	case circle:
		pf("%#v is circle\n", val)
		pf("volume = %.3f\n", val.volume())
	case rectangle:
		pf("%#v is rectangle\n", val)
		pf("diagonal = %.3f\n", val.diagonal())
	}

}
