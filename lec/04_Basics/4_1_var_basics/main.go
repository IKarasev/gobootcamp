package main

import "fmt"

func main() {
	// int var declaration and init
	var age int = 30
	fmt.Println("Given age", age)

	// string
	var name = "Ilya"
	fmt.Println("Given name", name)

	// not used var (caution)
	var notused = "Nonono"
	_ = notused

	// short decl and init
	message := "Some message!"
	_ = message

	// multiple decl
	car, model, cost := "Toyota", "HL", 40000
	fmt.Printf("Car %s %s: %d", car, model, cost)

	// another way to declare multiple vars
	var (
		firstName string
		married   bool
		salary    float64
	)
	fmt.Println(firstName, married, salary)
	var a, b, c int
	b, c = 3, 4
	fmt.Println(a, b, c)

	// types transl
	a1 := 4
	b1 := 5.1
	a1 = int(b1)
	fmt.Println(a1, b1)
}
