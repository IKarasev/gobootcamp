package main

import "fmt"

func main() {
	x := 10.10

	fmt.Println(&x)

	ptr := &x

	fmt.Printf("%T = %p\n", ptr, ptr)
	fmt.Printf("val at %p = %v", ptr, *ptr)
}
