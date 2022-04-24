package main

import "fmt"

func main() {
	fmt.Println("Array intro!")
	// init
	var numbers [4]int

	// printing
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n\n", a1)

	var a2 = [4]int{-10, 2, 10} // not all elems must be init. not init elems will get default val
	fmt.Printf("%#v\n\n", a2)

	// autodetect size of erray
	a3 := [...]string{"a", "b", "c", "d", "e", "f"}
	fmt.Printf("%#v\n\n", a3)
}
