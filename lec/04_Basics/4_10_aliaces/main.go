package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aliaces!!")
	var a uint8 = 8
	var b byte // byte is aliace for unit8
	b = a      // good, as b of aliace type of a
	fmt.Println(b)

	type second = uint // declare new aliace
	var hour second = 3600
	fmt.Printf("%d seconds is %d hours\n", hour, hour/60)

}
