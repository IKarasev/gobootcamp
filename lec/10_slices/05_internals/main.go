package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	fmt.Println("Slices internals!")
	fmt.Println(strings.Repeat("-", 10))

	// when creating slice go creates backing array
	// slice header - is a inner go structure to describe slice
	// slice elems:
	//	address - pointer to backing array start
	//	length - len of slice - returned by len()
	//	capacity - size of backing array - returned by cap()

	// nil slice - doesn't have header

	s1 := []int{10, 20, 30, 40, 50}
	fmt.Printf("s1: %#v\n", s1)

	s2, s3 := s1[:2], s1[1:3]
	fmt.Println("s2, s3 := s1[:2], s1[1:3]")
	fmt.Printf("s2: %#v\n", s2)
	fmt.Printf("s3: %#v\n", s3)

	fmt.Println(strings.Repeat("-", 10))

	// s1,s2,s3 - share same backing array
	// thus changing one of them changes all
	s3[1] = 666
	fmt.Printf("s3[1] = 666 => %#v\n", s3)
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)

	fmt.Println(strings.Repeat("-", 10))

	// if slice an array
	arr1 := [5]int{1, 2, 3, 4, 5}
	sl1, sl2 := arr1[:2], arr1[1:3]
	fmt.Printf("arr1: %#v\n", arr1)
	fmt.Printf("sl1: %#v\n", sl1)
	fmt.Printf("sl2: %#v\n\n", sl2)

	arr1[1] = 99
	fmt.Println("arr1[1] = 99")
	fmt.Printf("arr1: %#v\n", arr1)
	fmt.Printf("sl1: %#v\n", sl1)
	fmt.Printf("sl2: %#v\n", sl2)

	// arr1 - is backed array for new slices. Changes to array change all it slices

	fmt.Println(strings.Repeat("-", 10))

	cars := []string{"Toyota", "Volvo", "Nissan", "Lexus", "Ford"}
	newCars := []string{}

	newCars = append(newCars, cars[:2]...) // its a new slice with new backed array

	fmt.Printf("cars = %#v\n", cars)
	fmt.Printf("newCars = %#v\n", newCars)

	cars[0] = "Audi"
	fmt.Println("cars[0] = \"Audi\"")
	fmt.Printf("cars = %#v\n", cars)
	fmt.Printf("newCars = %#v\n", newCars)

	fmt.Println(strings.Repeat("-", 10))

	// len and cap
	s4 := []int{10, 20, 30, 40, 50}
	s5 := s4[0:3] // shares backed array
	fmt.Printf("s4 = %#v\n", s4)
	fmt.Printf("s5 = %#v\n", s5)
	fmt.Printf("len(s5) = %d\n", len(s5)) // len us 3
	fmt.Printf("cap(s5) = %d\n", cap(s5)) // cap is 5, number of elemets in backing array after
	// poiter to slice start in backing array

	s5 = s4[2:] //30 40 50
	fmt.Printf("s5 = %#v\n", s5)
	fmt.Printf("len(s5) = %d\n", len(s5)) // len is 3
	fmt.Printf("cap(s5) = %d\n", cap(s5)) // cap is 3, as there 3 elems in backed array  fter slice start

	fmt.Println(strings.Repeat("-", 10))

	// printing start address of slice
	fmt.Printf("s4 start = %p\n", &s4)
	fmt.Printf("s5 start = %p\n", &s5)

	fmt.Println(strings.Repeat("-", 10))

	// memory size of var
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("array = %#v\n", a1)
	fmt.Printf("slice = %#v\n", a2)
	fmt.Printf("array size, bytes = %d\n", unsafe.Sizeof(a1)) // => 40
	fmt.Printf("slice size, bytes = %d\n", unsafe.Sizeof(a2)) // => 24
	// aray size > slice size
}
