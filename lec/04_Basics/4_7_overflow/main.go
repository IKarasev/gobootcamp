package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("OVERFLOW!!")

	var x uint8 = 255 // max unit8 value
	x++               //overflow, no error, as it not at compile time
	fmt.Println(x)

	//a := int8(255 + 1) // error at compile time

	var b int8 = 127
	fmt.Printf("b: %d\n", b+1) // -> -128, overflow goes to min val for type
	b = -128
	b--
	fmt.Printf("b: %d\n", b) // -> 127, neg overflow goes to max val for type

	var f float32 = math.MaxFloat32 // -> 340282346638528859811704183484516925440.000000
	fmt.Printf("f: %f\n", f)

	f *= 2.0       // float32 overflow
	fmt.Println(f) // -> +Inf

	//const i int8 = 300 // error at compile: cannot use 300 (untyped int constant) as int8 value in constant declaration (overflows)
}
