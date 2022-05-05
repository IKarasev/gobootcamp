package main

import "fmt"

func main() {
	var m1 map[int]int
	fmt.Printf("%T: %#v\n", m1, m1)

	m2 := map[int]string{
		1: "A",
		2: "B",
	}

	m2[10] = "Abba"
	fmt.Println(m2[1])
	fmt.Println(m2[50])
}
