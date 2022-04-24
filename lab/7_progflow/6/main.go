package main

import (
	f "fmt"
)

func main() {
	age := -9
	switch {
	case age < 0 || age > 100:
		f.Println("Invalid age")
	case age < 18:
		f.Println("Minor")
	case age == 18:
		f.Println("Just major")
	default:
		f.Println("Major")
	}

}
