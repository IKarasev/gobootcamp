package main

import "fmt"

func main() {
	name := "Ilya"
	fmt.Println("Name: ", name)

	a, b := 3, 4
	fmt.Println("a + b = ", a+b)

	// formated print
	// string
	fmt.Printf("Name is %s \n", name)
	// int
	fmt.Printf("Age is %d \n", 30)
	// float
	fmt.Printf("pi = %f \n", 3.14159)
	fmt.Printf("pi = %.3f \n", 3.14159)
	// quoted str
	fmt.Printf("Pizza %q \n", "Margarita")
	// any type
	fmt.Printf("%v %v %v %v \n", "Vars", true, 4, 3.14)
	// print type
	fmt.Printf("%T \n", false)
	fmt.Printf("%T \n", 3.14)
	fmt.Printf("%T \n", "aaa")
	// bool
	fmt.Printf("%t \n", true)
	// base 2
	fmt.Printf("%b \n", 4)
	// base 2 with additional bits
	fmt.Printf("%08b \n", 4)
	// hex
	fmt.Printf("%x \n", 123)
}
