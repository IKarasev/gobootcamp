package main

import "fmt"

func main() {
	fmt.Println("Conditions!")
	age, mature := 20, false

	if age > 17 {
		mature = true
	} else {
		mature = false
	}

	fmt.Println(mature)
}
