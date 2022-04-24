package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch!")

	lang := "Golang"
	switch lang {
	case "Python":
		fmt.Println("Wrong")
		// break statement is outomatic
	case "Go", "Golang": // equivalent to || statement
		fmt.Println("Right! It's Golang!")
	default: // If no cases match
		fmt.Println("All lang is good!")
	}

	n := 5
	switch true { // used to switch over any bool cases. Just switch {} can be used
	case n%2 == 0:
		fmt.Println("Even number")
	case n%2 != 0: // in this case default for odd can be used
		fmt.Println("Odd number")
	default:
		fmt.Println("Other number")
	}

	hour := time.Now().Hour()
	fmt.Printf("Current hour: %v\n", hour)

	switch { // equals to switch true {}
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
