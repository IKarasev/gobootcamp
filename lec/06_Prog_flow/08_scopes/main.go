package main

import ( // file scope
	"fmt"
)

const pi = 3.14 // package scope

func main() { // package scope
	fmt.Println("Scopes!")

	x := 10 // local (block) scope
	fmt.Println(x, pi)
}
