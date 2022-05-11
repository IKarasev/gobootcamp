package main

import "fmt"

func swapFloat(x *float64, y *float64) {
	*x, *y = *y, *x
}

func main() {
	x, y := 5.5, 8.8
	swapFloat(&x, &y)
	fmt.Printf("x = %v\ny = %v\n", x, y)
}
