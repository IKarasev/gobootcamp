package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Functions: intro!"

// function declaration:
// func (reciver) name(params)(returns)  {
// 		...code
// }

func f1() {
	fmt.Println("First function!")
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// call func with no return
	f1()
}
