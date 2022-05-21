package main

import (
	"fmt"
	"gostudy/numbers" // importing our module
)

func main() {

	var n uint = 13

	// call func from module
	fmt.Printf("%d is even: %t\n", n, numbers.Even(n))

	n = 20
	fmt.Printf("%d is even: %t\n", n, numbers.Even(n))

	var intNum int = 32

	// call func from modele numbers, but from different file
	fmt.Printf("%d is prime: %t\n", intNum, numbers.IsPrime(intNum))

	intNum = 7
	fmt.Printf("%d is prime: %t\n", intNum, numbers.IsPrime(intNum))

	// try to call not global func from our package
	// numbers.notGlobalFunc() // => error: undefined: numbers.notGlobalFunc
	// can't access non global funcs

	// call package func that uses private funcs
	fmt.Printf("%d is odd and prime: %t\n", intNum, numbers.OddAndPrime(intNum))
}
