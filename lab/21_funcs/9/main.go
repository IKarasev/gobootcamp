package main

import (
	"fmt"
)

func f1(n int) {
	fmt.Println(n)
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	a(5)
	a(10)

}
