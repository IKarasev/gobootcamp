package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("NUMBER CONVERSION!!")

	x := 3
	y := 4.3
	fmt.Printf("x: %T\n", x)
	fmt.Printf("y: %T\n", y)

	//x = x * y // error, mismatched types int and float64
	x = x * int(y)
	fmt.Printf("x = %d\n", x) // -> 12

	x = 2
	//y = y * x // error, mismatched types float64 and int
	y = y * float64(x)
	fmt.Printf("y = %.3f", y) // -> 8.6

	var a = 5
	var b int64 = 10
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)

	//a = b // error
	a = int(b) // good
	fmt.Printf("a =  %d\n", a)

	// num to str and reverse
	fmt.Println("-- Num to str and rev --")

	s := string(99)
	fmt.Printf("s = %s\n", s) // -> c, 99 is ASCII for "c"

	//s := string(22.4) // error
	var s1 = fmt.Sprintf("%.2f", 22.44)
	fmt.Printf("Float string s1 = %s\n", s1)

	s = fmt.Sprintf("%d", 1234)
	fmt.Printf("Int string s = %s\n", s)

	// convert str to num
	s = "3.1415"
	fmt.Printf("s: %T\n", s)
	var f1, err = strconv.ParseFloat(s, 64) // f1 - result of conv, err - error message
	fmt.Printf("f1 = %f\n", f1)
	fmt.Printf("err: %v\n", err)
	f1, err = strconv.ParseFloat("asdsd", 64)
	fmt.Printf("f1 = %f\n", f1)
	fmt.Printf("err: %v\n", err)

	// str to int
	i, err := strconv.Atoi("55")
	fmt.Printf("i = %d\n", i)
	fmt.Printf("err: %v\n", err)

	// int to ascii
	s2 := strconv.Itoa(43)
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s2 = %q\n", s2)

}
