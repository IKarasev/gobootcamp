package main

import (
	"fmt"
)

func f1(n uint) (f int, s int) {
	s = 0
	f = 1

	for i := 1; i <= int(n); i++ {
		s += i
		f *= i
	}

	return
}

func main() {
	f, s := f1(3)
	fmt.Println("funtorial 3 =", f)
	fmt.Println("sum till 3 =", s)

	f, s = f1(5)
	fmt.Println("funtorial 5 =", f)
	fmt.Println("sum till 5 =", s)
}
