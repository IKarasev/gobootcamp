package main

import "fmt"

func main() {
	fmt.Println("OPERATORES!!")

	a, b := 4, 2
	r := (a + b) / (a - b) * 2 // = 6
	fmt.Println("(a +b)/(a-b)*2 = ", r)
	r = 9 % a // = 1
	fmt.Printf("9 %% %d = %v\n", a, r)

	a, b = 2, 3

	// incr assignment
	a += b
	fmt.Println(a)
	// decr assignment
	b -= 2
	fmt.Println(b)
	// mult
	b *= 10
	fmt.Println(b)
	// div
	b /= 5
	fmt.Println(b)
	// mod
	a %= 3
	fmt.Println(a)

	// incr
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// logical operators
	fmt.Println("logic operators:")
	fmt.Println(5 == 2)
	fmt.Println("aaa" == "aaa")
	fmt.Println(5 > 4)
	fmt.Println(5 >= 5)
	fmt.Println(5 < 4)
	fmt.Println(5 <= 5)
	fmt.Println(true != false)
}
