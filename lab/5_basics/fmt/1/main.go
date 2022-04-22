package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("%d %f %s %v\n", x, y, z, score)
	fmt.Printf("\"%s\"\n", z)
	fmt.Printf("%T %T \n", y, score)
}
