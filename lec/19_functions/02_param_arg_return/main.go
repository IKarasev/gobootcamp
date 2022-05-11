package main

import (
	"fmt"
	"math"
	"strings"
)

const lecHeader string = "Functions: params, args, return!"

// simple func
func f1() {
	fmt.Println("First function!")
}

// func with params
func f2(a int, b int) {
	fmt.Println(">>f2():")
	fmt.Printf("param a = %d\nparam b = %d\n", a, b)
	fmt.Println("a + b =", a+b)
}

// short param declarations
func f3(a, b, c int, d, e, f string) {
	fmt.Println(">>f3():")
	fmt.Println("int params:", a, b, c)
	fmt.Println("str params:", d, e, f)
}

// func with return   \/ note return type in func
func f4(a, b float64) float64 {
	fmt.Println(">>f4():")
	return math.Pow(a, b)
}

// func with multiple val return, if error is returned - it should be last return param
func f5(a, b int) (int, int) {
	fmt.Println(">>f5():")
	return a + b, a * b
}

// named return values
func f6(a, b int) (s int) {
	fmt.Println(">>f6():")
	fmt.Println("s =", s)

	s = a + b
	return // note: with named return no need to specify return var
	// also called Naked return
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// call func with no return
	f1()

	// call func with params
	f2(10, 15)
	f3(1, 2, 3, "go", "func", "learn")

	// call func with return
	pln("2.2 ^ 3.3 = ", f4(2.2, 3.3))
	r := f4(5.1, 2.5)
	pln("r = ", r)

	// call func with mult return
	r1, r2 := f5(2, 3)
	pf("2+3=%d\n2*3=%d\n", r1, r2)

	pln("5 + 8 =", f6(5, 8))
}
