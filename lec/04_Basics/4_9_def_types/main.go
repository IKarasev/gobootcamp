package main

import (
	"fmt"
)

type km float64
type mile float64

func main() {
	fmt.Println("DEFINED (NAMED) TYPES!!")

	// declaration
	type age int        // int is underlying type
	type oldAge age     // int is underlying type
	type veryOldAge age // int is underlying type

	type speed uint
	var s1 speed = 10
	fmt.Printf("s1: %T = %v\n", s1, s1)

	var s2 speed = 20
	fmt.Println(s1 + s2)

	var x uint
	x = uint(s1)
	fmt.Println(x)

	//var s3 speed = x //error, (variable of type uint) as type speed in variable declaration
	var s3 speed = speed(x)
	fmt.Println(s3)

	var ptL km = 465
	var distanceMiles mile
	distanceMiles = mile(ptL) / 0.621
	fmt.Println(distanceMiles)

}
