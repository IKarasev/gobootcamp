package main

import "fmt"

func main() {
	const (
		secPerDay int = 86400
		daysYear  int = 365
	)

	fmt.Printf("%d\n", secPerDay*daysYear)
}
