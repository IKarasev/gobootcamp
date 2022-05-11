package main

import (
	"fmt"
	"math"
)

func cube(x float64) float64 {
	return math.Pow(x, 3)
}

func main() {
	fmt.Printf("2^3 = %.3f\n", cube(2))
	fmt.Printf("5.5^3 = %.3f\n", cube(5.5))
}
