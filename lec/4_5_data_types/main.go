package main

import "fmt"

func main() {
	fmt.Println("Data Types")

	// INT TYPES
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	// FLOAT
	var f1, f2, f3 float64 = 2.3, -.2, 5.
	fmt.Printf("%.2f %.2f %.2f\n", f1, f2, f3)

	// byte = unit8
	// rune = int32 (used as char)
	var r rune = 'f'
	fmt.Printf("f is %T :: %d is decimal ASCII code\n", r, r)
	fmt.Printf("		%x is hex ASCII code\n", r)

	// BOOL
	var b bool = true
	fmt.Printf("%T\n", b)

	// STRING
	var s string = "AAAAA"
	fmt.Printf("%T\n", s)

	// ARRAY
	var numbers = [4]int{4, 5, -9, 122}
	fmt.Println("ARRAYS")
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)

	// SLICE
	var animal = []string{"dog", "degu", "rino"}
	fmt.Println("SLICE")
	fmt.Printf("%T\n", animal)
	fmt.Println(animal)

	// MAP ~ dict in Python
	curencies := map[string]float64{
		"USD": 80.10,
		"EUR": 95.22,
		"RUB": 1,
	}
	fmt.Printf("%T\n", curencies)
	fmt.Println(curencies)

	// STRUCT ~ structure in c++
	type Cat struct {
		name  string
		age   int
		breed string
		male  bool
	}

	var Basya Cat
	fmt.Printf("%T\n", Basya)
	fmt.Println(Basya)
	Basya.name = "Basya"
	Basya.age = 4
	Basya.breed = "British"
	Basya.male = false
	fmt.Println(Basya)

	// POINTER
	var x int = 2
	ptr := &x
	fmt.Printf("ptr type [%T] and value [%v]\n", ptr, ptr)

	// FUNCTION
	fmt.Printf("%T\n", f)
}

func f() {

}
