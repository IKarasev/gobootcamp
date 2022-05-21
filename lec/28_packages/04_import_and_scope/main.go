package main

import "fmt"

// using "go run *.go" we can run all .go files in dir
func main() {
	fmt.Println("Import and Scope!")
	sayHello()

	fmt.Println("H2O boiling point in Celsius:", H2OBoilingPoint)
	fmt.Println("H2O boiling point in Farenheit:", toFarenheit(H2OBoilingPoint))
}
