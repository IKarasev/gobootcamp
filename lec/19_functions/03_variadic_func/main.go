package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Functions: Variadic functions!"

// creating variadic func
func f1(a ...int) {
	fmt.Println(">>f1():")
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

// do something
func f2(a ...int) {
	fmt.Println(">>f2():")
	a[0] = 50
}

// with return
func f3(a ...int) (s int, p int) {
	fmt.Println(">>f3():")
	p = 1
	for _, v := range a {
		s = s + v
		p = p * v
	}

	return
}

// mix periodic and variadic params
func f4(age int, names ...string) string {
	fmt.Println(">>f4():")
	fullName := strings.Join(names, " ")
	result := fmt.Sprintf("Age: %d, Name: %q", age, fullName)
	return result
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	f1()
	f1(1)
	f1(1, 2, 3, 4)
	a := []int{10, 20, 30}
	f1(a...)

	pln(strings.Repeat("-", 10))

	f2(a...)
	pln(a)

	pln(strings.Repeat("-", 10))
	pln(f3(1, 2, 3, 4, 5))

	pln(strings.Repeat("-", 10))
	pln(f4(25, "John", "Malkovich", "Junior"))
}
