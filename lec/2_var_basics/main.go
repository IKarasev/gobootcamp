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
}
