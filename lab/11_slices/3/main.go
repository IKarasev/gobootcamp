package main

import "fmt"

func main() {
	s := make([]float64, 3)
	s = append(s, 10.1)
	s = append(s, 4.1, 5.5, 6.6)

	fmt.Println(s)

	n := make([]float64, 2)
	n = append(n, s...)
	fmt.Println(n)
}
