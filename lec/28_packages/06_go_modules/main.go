package main

import (
	"fmt"

	"github.com/IKarasev/gotest/mygeometry"
	"github.com/IKarasev/gotest/mymath"
)

func main() {
	fmt.Println("7 is even:", mymath.Even(7))
	fmt.Println("7 is is odd:", mymath.Odd(7))
	fmt.Println("7 is prime:", mymath.IsPrime(7))

	fmt.Println("Perimeter of Circle r = 14:", mygeometry.CirclePerimeter(14.))
	fmt.Println("rectangle 5x8 diagonal:", mygeometry.RecDiagonal(5., 8.))
}
