package main

import "fmt"

const H2OBoilingPoint = 100

func sayHello() {
	fmt.Println("My Hello!")
}

func toFarenheit(t float64) float64 {
	return t*1.8 + 32
}
