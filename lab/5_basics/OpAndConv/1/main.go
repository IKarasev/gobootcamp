package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = 3.2

	var i1 = float64(i)
	var f1 = int(f)

	fmt.Printf("i type %T val %v\n", i1, i1)
	fmt.Printf("f type %T val %v\n", f1, f1)
}
